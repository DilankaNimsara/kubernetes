# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Install CA certificates and Git (for module fetching if needed)
RUN apk update && apk add --no-cache ca-certificates git

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o kube-server .

# Stage 2: Create a minimal final image
FROM alpine:latest
WORKDIR /usr/local/bin

# Install CA certificates (required for HTTPS in Go apps)
RUN apk --no-cache add ca-certificates

# Copy binary from builder stage
COPY --from=builder /app/kube-server .

# Expose the application's port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./kube-server"]
