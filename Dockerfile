# --- Build Stage ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# Download all dependencies.
# Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the api server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/server

# Build the worker
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o worker ./cmd/worker


# --- Final Stage ---
FROM alpine:latest

WORKDIR /app

# Copy the binaries from the builder stage
COPY --from=builder /app/api .
COPY --from=builder /app/worker .

# Expose port for the api
EXPOSE 8080

# The command will be specified in the docker-compose.yaml
# For example: CMD ["./api"] or CMD ["./worker"]
