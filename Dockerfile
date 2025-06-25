# Build da aplicação Go
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app ./cmd/api

# Imagem final com NATS + app + supervisord
FROM alpine:latest

RUN apk add --no-cache curl ca-certificates supervisor

# Baixa e instala NATS
RUN curl -L https://github.com/nats-io/nats-server/releases/download/v2.10.11/nats-server-v2.10.11-linux-amd64.tar.gz \
    | tar xz && \
    mv nats-server-v2.10.11-linux-amd64/nats-server /usr/local/bin/ && \
    rm -rf nats-server-v2.10.11-linux-amd64

# Cria diretórios para supervisord
RUN mkdir -p /var/log /var/run

# Copia app Go compilado
COPY --from=builder /app/app /usr/local/bin/app

# Copia configuração do supervisor
COPY supervisord.conf /etc/supervisord.conf

EXPOSE 8080 4222

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
