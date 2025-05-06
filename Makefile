build:
	go build -o ./bin/cli ./cmd/api/main.go

run: build
	go run ./cmd/api

test:
	go test ./... -v
	@echo "Running tests..."

lint:
	golangci-lint run ./...