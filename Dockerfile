FROM golang:1.24.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

WORKDIR /app/cmd/server

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]