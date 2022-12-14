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

// HandleCreateTaskRequest handles createTask operation.
//
// POST /tasks
func (s *Server) handleCreateTaskRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createTask"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateTask",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()
	request, err := decodeCreateTaskRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreateTask(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateTaskResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	span.SetStatus(codes.Ok, "Ok")
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleDeleteTaskRequest handles deleteTask operation.
//
// DELETE /tasks/{id}
func (s *Server) handleDeleteTaskRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteTask"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteTask",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()
	params, err := decodeDeleteTaskParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeleteTask(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeleteTaskResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	span.SetStatus(codes.Ok, "Ok")
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleListTaskRequest handles listTask operation.
//
// GET /tasks
func (s *Server) handleListTaskRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listTask"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListTask",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()
	params, err := decodeListTaskParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListTask(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListTaskResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	span.SetStatus(codes.Ok, "Ok")
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleListUserTasksRequest handles listUserTasks operation.
//
// GET /users/{id}/tasks
func (s *Server) handleListUserTasksRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listUserTasks"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListUserTasks",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()
	params, err := decodeListUserTasksParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListUserTasks(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListUserTasksResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	span.SetStatus(codes.Ok, "Ok")
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleReadTaskRequest handles readTask operation.
//
// GET /tasks/{id}
func (s *Server) handleReadTaskRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readTask"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadTask",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()
	params, err := decodeReadTaskParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadTask(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadTaskResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	span.SetStatus(codes.Ok, "Ok")
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleReadTaskAssignedRequest handles readTaskAssigned operation.
//
// GET /tasks/{id}/assigned
func (s *Server) handleReadTaskAssignedRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readTaskAssigned"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadTaskAssigned",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()
	params, err := decodeReadTaskAssignedParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadTaskAssigned(ctx, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadTaskAssignedResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	span.SetStatus(codes.Ok, "Ok")
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleUpdateTaskRequest handles updateTask operation.
//
// PATCH /tasks/{id}
func (s *Server) handleUpdateTaskRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateTask"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdateTask",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()
	params, err := decodeUpdateTaskParams(args, r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeUpdateTaskRequest(r, span)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "BadRequest")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.UpdateTask(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeUpdateTaskResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	span.SetStatus(codes.Ok, "Ok")
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}
