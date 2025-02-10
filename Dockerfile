FROM golang:1.21-alpine AS builder
WORKDIR /app
# Add git for fetching dependencies
RUN apk add --no-cache git

COPY . .
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o exporter ./cmd/exporter

FROM alpine:3.17

RUN apk add --no-cache ipmitool
COPY --from=builder /app/exporter /exporter

EXPOSE 9290
ENTRYPOINT ["/exporter"]