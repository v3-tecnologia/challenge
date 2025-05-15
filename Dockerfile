# Use the official Golang image as a build stage
FROM golang:1.24.3 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Set environment variables for Go build
ENV GOARCH=amd64
ENV GOOS=linux

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go app from the new entrypoint path
RUN go build -o main ./cmd/server

# Start a new stage with Debian
FROM debian:12

# Set the Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x /app/main

# Command to run the executable
CMD ["/app/main"]
