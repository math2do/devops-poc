# Start with the official Golang base image
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go appl
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/server ./cmd/server

# Use a minimal base image for the final build
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/bin/server .

# Expose port (adjust if your service uses a different port)
EXPOSE 8080

# Command to run the executable
CMD ["./server"]
