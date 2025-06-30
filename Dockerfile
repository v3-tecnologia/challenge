FROM golang:alpine AS build
RUN apk add alpine-sdk
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=1
RUN go build

FROM scratch
COPY --from=build /app/mychallenge /app
EXPOSE 3000
ENV PORT=3000
CMD [ "/app" ]
