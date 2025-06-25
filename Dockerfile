FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app ./cmd/api

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y supervisor curl ca-certificates && rm -rf /var/lib/apt/lists/*

# Instalar NATS
RUN curl -L https://github.com/nats-io/nats-server/releases/download/v2.10.11/nats-server-v2.10.11-linux-amd64.tar.gz \
    | tar xz && \
    mv nats-server-v2.10.11-linux-amd64/nats-server /usr/local/bin/ && \
    rm -rf nats-server-v2.10.11-linux-amd64

COPY --from=builder /app/app /usr/local/bin/app
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 8080 4222

CMD ["supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
