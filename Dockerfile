# Stage: build
FROM golang:1.24.5 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/mini-monitoring-app

# Stage: runtime (alpine)
FROM alpine:3.19

WORKDIR /app

# copy binary from builder
COPY --from=builder /app/mini-monitoring-app /app/mini-monitoring-app

# (optional) run as non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown appuser:appgroup /app/mini-monitoring-app
USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/mini-monitoring-app"]
