build:
	@echo "Building..."
	@go build -v -o bin/challenge cmd/api/main.go

run:
	@go run cmd/api/main.go

test:
	@go test -v ./...
