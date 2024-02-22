package events

import (
	"context"
)

const (
	CUDTypeCreate    = "CREATE"
	CUDTypeUpdate    = "UPDATE"
	CUDTypeDelete    = "DELETE"
	CUDTypeBuissness = "BUISNESS"

	EventHeaderID   = "EVENT_ID"
	EventHeaderType = "EVENTE_TYPE"
)

type (
	EventConstraint interface {
		Event
	}

	Reader[E EventConstraint] interface {
		Read(ctx context.Context, e E) (E, error)
	}

	Writer[E EventConstraint] interface {
		Write(ctx context.Context, e E) (E, error)
	}

	CUDType string

	Event struct {
		EventType CUDType `json:"eventType"`
	}
)
