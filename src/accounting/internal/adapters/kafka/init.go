package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/caarlos0/env"
	"github.com/google/uuid"
	"github.com/popoffvg/async-arch/accounting/internal/ent"
	"github.com/popoffvg/async-arch/common/pkg/logger"
	"github.com/segmentio/kafka-go"
	"time"
)

const (
	topicAssign = "assign"
	topicUsers  = "users"
	topicTasks  = "tasks"

	groupID = "accounting"
)

type (
	Config struct {
		Address string `env:"KAFKA_ADDRESS" envDefault:"localhost:9093"`
	}

	Adapter struct {
		orm *ent.Client

		rTasks  *kafka.Reader
		rAssign *kafka.Reader
		rUsers  *kafka.Reader

		log logger.Logger

		cancels []context.CancelFunc
	}
)

func New(
	cfg *Config,
	log logger.Logger,
	client *ent.Client,
) *Adapter {
	a := new(Adapter)
	a.log = log
	a.orm = client

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	a.rUsers = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.Address},
		GroupID: groupID,
		Topic:   topicUsers,
		Dialer:  dialer,
	})

	a.rAssign = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.Address},
		GroupID: groupID,
		Topic:   topicAssign,
		Dialer:  dialer,
	})

	a.rTasks = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.Address},
		GroupID: groupID,
		Topic:   topicTasks,
		Dialer:  dialer,
	})

	return a
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("read kafka config failed: %w", err)
	}

	return &cfg, nil
}

func writeMessage(ctx context.Context, w *kafka.Writer, msg any) (err error) {
	var payload []byte
	if payload, err = json.Marshal(msg); err != nil {
		return err
	}

	return retry.Do(func() error {
		return w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(uuid.New().String()),
			Value: payload,
		})

	},
		retry.Attempts(3),
		retry.Delay(3*time.Second),
	)
}

func readMessage[T any](ctx context.Context, r *kafka.Reader) (context.Context, *T, error) {
	msg, err := r.ReadMessage(ctx) // TODO: read with immidetly commit for simplest solution
	if err != nil {
		return ctx, nil, err
	}

	var payload T
	if err := json.Unmarshal(msg.Value, &payload); err != nil {
		return ctx, nil, err
	}

	ctx = context.WithValue(ctx, EventIDKey, msg.Key)
	return ctx, &payload, nil
}

func startRead(a *Adapter) {
	ctx, cancel := context.WithCancel(context.Background())
	a.cancels = append(a.cancels, cancel)
	go a.ReadUserCUD(ctx)
	a.cancels = append(a.cancels, func() {
		a.rTasks.Close()
	})
	go a.ReadTaskCUD(ctx)
	a.cancels = append(a.cancels, func() {
		a.rAssign.Close()
	})
	go a.ReadTaskAssign(ctx)
	a.cancels = append(a.cancels, func() {
		a.rUsers.Close()
	})
}
