set -e  # Stop the script on any error

echo "🔁 Cleaning up previous containers..."
docker compose -f docker-compose.dev.test.yml down -v

echo "🚀 Running integration tests..."
docker compose -f docker-compose.dev.test.yml up --build --abort-on-container-exit
integration_status=$?

echo "🧹 Shutting down containers..."
docker compose -f docker-compose.dev.test.yml down -v

if [ $integration_status -ne 0 ]; then
  echo "❌ Integration tests failed!"
  exit $integration_status
fi

echo "🧪 Running unit tests..."
go test ./internal/...
unit_status=$?

if [ $unit_status -ne 0 ]; then
  echo "❌ Unit tests failed!"
  exit $unit_status
fi

echo "✅ All tests passed successfully!"
