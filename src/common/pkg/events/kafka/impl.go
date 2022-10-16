package kafka

import (
	"context"
	"github.com/ogen-go/ogen/json"
	"github.com/popoffvg/async-arch/common/pkg/events"
	"github.com/segmentio/kafka-go"
)

type (
	ReaderWrap[E events.EventConstraint] struct {
		r *kafka.Reader
	}

	WriterWrap struct {
		w *kafka.Writer
	}
)

func (r *ReaderWrap[E]) Read(ctx context.Context) (E, error) {
	var e E

	msg, err := r.r.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(msg.Value, &e); err != nil {
		return nil, err
	}

	headers := getHeaders(msg)
	e.Timestamp = msg.Time
	if v, ok := headers[events.EventHeaderID]; ok {
		e.EventID = string(v)
	}

	if v, ok := headers[events.EventHeaderType]; ok {
		e.EventType = events.CUDType(v)
	}

	return e, nil
}

func getHeaders(msg kafka.Message) map[string][]byte {
	hs := make(map[string][]byte, len(msg.Headers))
	for _, v := range msg.Headers {
		hs[v.Key] = v.Value
	}

	return hs
}
