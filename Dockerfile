# Dockerfile
FROM golang:1.24.4 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy

RUN go test ./internal/tests/...

# Build binário estático para Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM debian:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

CMD ["/app/main"]
