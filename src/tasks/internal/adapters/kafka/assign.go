package kafka

import (
	"context"
	"time"
)

type (
	Message struct {
		Date time.Time `json:"date"`
	}
)

func (a *Adapter) PlanAssign(ctx context.Context) error {
	return writeMessage(ctx, a.wAssign, Message{
		Date: time.Now(),
	})
}

func (a *Adapter) ProduceAssign(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:

		}
		_, err := readMessage[Message](ctx, a.rAssign)
		if err != nil {
			a.log.Errorf("read message from topic %s failed: %s", topicAssign, err.Error())
			continue
		}

		err = a.tasks.AssignTasks(ctx)
		if err != nil {
			a.log.Errorf("assign task failed: %s", err.Error())
		}

	}
}
