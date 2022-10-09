package tasks

import "github.com/popoffvg/async-arch/tasks/ent"

type Service struct {
	client *ent.Client
}

func New(c *ent.Client) (*Service, error) {
	return &Service{client: c}, nil
}
