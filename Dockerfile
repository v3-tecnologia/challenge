# Etapa de build
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copia os arquivos Go para dentro da imagem
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila o binário principal
RUN go build -o app ./cmd/api

# Etapa final
FROM alpine:latest

WORKDIR /app

# Copia apenas o binário para a imagem final
COPY --from=builder /app/app .

EXPOSE 10000

CMD ["./app"]
