# Stage 1: Build
FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the binary
RUN go build -o mini-monitoring-app

# Stage 2: Run (lighter image)
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/mini-monitoring-app .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./mini-monitoring-app"]
