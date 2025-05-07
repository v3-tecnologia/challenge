FROM golang:1.24

WORKDIR /cmd/api

# Instalar ferramentas para verificação de saúde
RUN apt-get update && apt-get install -y postgresql-client

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

EXPOSE 8080

# Criar script de entrypoint
RUN echo '#!/bin/sh\n\
echo "Waiting for PostgreSQL to start..."\n\
until pg_isready -h db -U postgres; do\n\
  echo "PostgreSQL not ready yet - waiting..."\n\
  sleep 3\n\
done\n\
echo "PostgreSQL is up - starting application"\n\
env | grep DB_\n\
./bin/cli\n' > /entrypoint.sh && chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]
