# Etapa 1: build da aplicação
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app ./cmd/api

# Etapa 2: imagem final com NATS + App
FROM alpine:latest

# Instalar dependências
RUN apk add --no-cache curl ca-certificates supervisor

# Instalar NATS
RUN curl -L https://github.com/nats-io/nats-server/releases/download/v2.10.11/nats-server-v2.10.11-linux-amd64.tar.gz \
    | tar xz && \
    mv nats-server-v2.10.11-linux-amd64/nats-server /usr/local/bin/ && \
    rm -rf nats-server-v2.10.11-linux-amd64

# Copiar app Go compilado
COPY --from=builder /app/app /usr/local/bin/app

# Copiar config do supervisor
COPY supervisord.conf /etc/supervisord.conf

EXPOSE 8080 4222

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
