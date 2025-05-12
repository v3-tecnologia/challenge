#!/bin/bash
echo "Running tests..."
go test -v ./v3/internal/tests/...
TEST_EXIT_CODE=$?
if [ $TEST_EXIT_CODE -ne 0 ]; then
    echo "Tests failed. Exiting."
    exit $TEST_EXIT_CODE
fi
echo "Tests passed. Starting application..."
exec go run cmd/main.go