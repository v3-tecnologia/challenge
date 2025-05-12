FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o telemetry-challenge-api ./cmd/api

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache curl postgresql-client bash && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate

COPY --from=builder /app/telemetry-challenge-api /app/
COPY --from=builder /app/migrations /app/migrations

RUN echo '#!/bin/bash\n\
set -e\n\
\n\
echo "Aguardando PostgreSQL..."\n\
until PGPASSWORD=postgres psql -h db -U postgres -d telemetry-challenge-database -c "SELECT 1" > /dev/null 2>&1; do\n\
  echo "PostgreSQL não está pronto - aguardando..."\n\
  sleep 2\n\
done\n\
\n\
echo "PostgreSQL pronto!"\n\
echo "Executando migrações..."\n\
migrate -path /app/migrations -database "$DB_URL" up\n\
\n\
echo "Iniciando aplicação..."\n\
exec /app/telemetry-challenge-api\n' > /app/entrypoint.sh && \
    chmod +x /app/entrypoint.sh

# Definir variáveis de ambiente
ENV DB_URL=postgresql://postgres:postgres@db:5432/telemetry-challenge-database?sslmode=disable
ENV MIGRATIONS_DIR=/app/migrations

# Executar o script de entrypoint
CMD ["/app/entrypoint.sh"]