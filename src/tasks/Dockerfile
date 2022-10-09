FROM golang:1.19-alpine as builder

ENV GOPROXY=direct       
ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
COPY go.* ./
RUN go build  -o /bin/app cmd/main.go

FROM alpine
COPY --from=builder /bin/app app

ENTRYPOINT [app]