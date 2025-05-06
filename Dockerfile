FROM golang:1.24.2-alpine

WORKDIR /app

RUN apk add --no-cache build-base

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o app ./main.go

EXPOSE 8080

CMD ["./app"]
