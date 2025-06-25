# Build stage para a API Golang
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/main.go

# Final stage
FROM alpine:latest

# Instalar dependências
RUN apk --no-cache add ca-certificates tzdata curl netcat-openbsd

# Baixar e instalar NATS Server
RUN curl -L https://github.com/nats-io/nats-server/releases/download/v2.10.7/nats-server-v2.10.7-linux-amd64.zip -o nats.zip && \
    unzip nats.zip && \
    mv nats-server-v2.10.7-linux-amd64/nats-server /usr/local/bin/ && \
    rm -rf nats* && \
    chmod +x /usr/local/bin/nats-server

WORKDIR /app

# Copy da API e script
COPY --from=builder /app/api .
COPY start.sh .
RUN chmod +x start.sh

# Expor porta da API (NATS roda internamente)
EXPOSE 8080

# Usar o script de inicialização
CMD ["./start.sh"]