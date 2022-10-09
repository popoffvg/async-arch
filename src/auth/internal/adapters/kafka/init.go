package kafka

import (
	"context"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/caarlos0/env"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/json"
	"github.com/popoffvg/async-arch/common/pkg/logger"
	"github.com/segmentio/kafka-go"
	"time"
)

const (
	topicUsers            = "users"
	topicUserRolesChanged = "user-roles-changed"
)

type (
	Config struct {
		Address string `env:"KAFKA_ADDRESS" envDefault:"localhost:9093"`
	}

	Adapter struct {
		wUsers            *kafka.Writer
		wUsersRoleChanged *kafka.Writer

		log logger.Logger

		cancels []context.CancelFunc
	}
)

func New(
	cfg *Config,
	log logger.Logger,
) (*Adapter, error) {
	a := new(Adapter)
	a.log = log.Named("kafka-adapter")

	_, err := kafka.DialLeader(context.Background(), "tcp", cfg.Address, topicUsers, 0)
	if err != nil {
		return nil, fmt.Errorf("create topic %s failed: %w", topicUsers, err)
	}
	_, err = kafka.DialLeader(context.Background(), "tcp", cfg.Address, topicUserRolesChanged, 0)
	if err != nil {
		return nil, fmt.Errorf("create topic %s failed: %w", topicUserRolesChanged, err)
	}

	a.wUsersRoleChanged = &kafka.Writer{
		Addr:     kafka.TCP(cfg.Address),
		Topic:    topicUserRolesChanged,
		Balancer: &kafka.LeastBytes{},
	}

	a.wUsers = &kafka.Writer{
		Addr:     kafka.TCP(cfg.Address),
		Topic:    topicUsers,
		Balancer: &kafka.LeastBytes{},
	}

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
