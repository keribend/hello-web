# --- Build stage ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app (replace 'main.go' with your entrypoint if different)
RUN go build -o hello-web cmd/hello-web/*.go

# --- Run stage ---
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder
COPY --from=builder /app/hello-web .

# Expose the port your app listens on (change if needed)
EXPOSE 8080

# Run the binary
CMD ["./hello-web"]