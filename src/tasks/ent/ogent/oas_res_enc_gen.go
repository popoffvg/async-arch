// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

func encodeCreateTaskResponse(response CreateTaskRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *TaskCreate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/tasks"+`: unexpected response type: %T`, response)
	}
}

func encodeDeleteTaskResponse(response DeleteTaskRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteTaskNoContent:
		w.WriteHeader(204)
		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/tasks/{id}"+`: unexpected response type: %T`, response)
	}
}

func encodeListTaskResponse(response ListTaskRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListTaskOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/tasks"+`: unexpected response type: %T`, response)
	}
}

func encodeListUserTasksResponse(response ListUserTasksRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListUserTasksOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/users/{id}/tasks"+`: unexpected response type: %T`, response)
	}
}

func encodeReadTaskResponse(response ReadTaskRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *TaskRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/tasks/{id}"+`: unexpected response type: %T`, response)
	}
}

func encodeReadTaskAssignedResponse(response ReadTaskAssignedRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *TaskAssignedRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/tasks/{id}/assigned"+`: unexpected response type: %T`, response)
	}
}

func encodeUpdateTaskResponse(response UpdateTaskRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *TaskUpdate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		e := jx.GetWriter()
		defer jx.PutWriter(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil
	default:
		return errors.Errorf("/tasks/{id}"+`: unexpected response type: %T`, response)
	}
}
