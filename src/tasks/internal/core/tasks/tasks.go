package tasks

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/popoffvg/async-arch/tasks/ent/task"
	"math/rand"
	"time"
)

func (s *Service) AssignTasks(ctx context.Context) error {
	users, err := s.client.User.Query().All(ctx)
	if err != nil {
		return fmt.Errorf("get useres for assinigning failed: %w", err)
	}

	tasks, err := s.client.Task.Query().
		Where(
			task.IsDone(false),
		).
		All(ctx)
	if err != nil {
		return fmt.Errorf("task assigning failed: %w", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var gerr *multierror.Error
	for _, v := range tasks {
		err = s.client.Task.UpdateOne(v).SetAssignedID(users[r.Intn(len(users)-1)].ID).Exec(ctx)
		if err != nil {
			gerr = multierror.Append(gerr, err)
		}
	}

	return gerr.ErrorOrNil()
}
