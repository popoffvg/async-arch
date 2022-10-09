package kafka

import (
	"context"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/caarlos0/env"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/json"
	"github.com/popoffvg/async-arch/common/pkg/logger"
	"github.com/popoffvg/async-arch/tasks/ent"
	"github.com/popoffvg/async-arch/tasks/internal/core/tasks"
	"github.com/segmentio/kafka-go"
	"time"
)

const (
	topicAssign = "assign"
	topicUsers  = "users"
)

type (
	Config struct {
		Address string `env:"KAFKA_ADDRESS" envDefault:"localhost:9092"`
	}

	Adapter struct {
		wAssign *kafka.Writer
		rAssign *kafka.Reader
		rUsers  *kafka.Reader

		tasks *tasks.Service
		log   logger.Logger
		orm   *ent.Client

		cancels []context.CancelFunc
	}
)

func New(
	cfg *Config,
	tasks *tasks.Service,
	log logger.Logger,
	client *ent.Client,
) (*Adapter, error) {
	a := new(Adapter)
	a.tasks = tasks
	a.log = log.Named("kafka-adapter")
	a.orm = client
	a.wAssign = &kafka.Writer{
		Addr:     kafka.TCP(cfg.Address),
		Topic:    topicAssign,
		Balancer: &kafka.LeastBytes{},
	}

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	a.rAssign = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.Address},
		GroupID: "",
		Topic:   topicAssign,
		Dialer:  dialer,
	})
	a.rUsers = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.Address},
		GroupID: "",
		Topic:   topicUsers,
		Dialer:  dialer,
	})

	return a, nil
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

func readMessage[T any](ctx context.Context, r *kafka.Reader) (*T, error) {
	msg, err := r.ReadMessage(ctx) // TODO: read with immidetly commit for simplest solution
	if err != nil {
		return nil, err
	}

	var payload T
	if err := json.Unmarshal(msg.Value, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func startRead(a *Adapter) {
	ctx, cancel := context.WithCancel(context.Background())
	a.cancels = append(a.cancels, cancel)
	go a.ProduceAssign(ctx)
}
