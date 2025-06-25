# Etapa 1: build da aplicação
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app ./cmd/api

# Etapa 2: imagem final
FROM alpine:latest

# Instalar dependências
RUN apk add --no-cache curl ca-certificates

# Instalar NATS
RUN curl -L https://github.com/nats-io/nats-server/releases/download/v2.10.11/nats-server-v2.10.11-linux-amd64.tar.gz \
    | tar xz && \
    mv nats-server-v2.10.11-linux-amd64/nats-server /usr/local/bin/ && \
    rm -rf nats-server-v2.10.11-linux-amd64

# Copiar app compilado e script de start
COPY --from=builder /app/app /usr/local/bin/app
COPY start.sh /usr/local/bin/start.sh
RUN chmod +x /usr/local/bin/start.sh

EXPOSE 8080 4222

CMD ["/usr/local/bin/start.sh"]
