FROM golang:1.24.3

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

# Copiar e dar permissão ao script de entrada
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 8080

# Usar o script como ponto de entrada
ENTRYPOINT ["./entrypoint.sh"]