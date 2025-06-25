# Etapa 2: imagem final com NATS + App
FROM alpine:latest

# Instalar dependências
RUN apk add --no-cache curl ca-certificates

# Instalar NATS
RUN curl -L https://github.com/nats-io/nats-server/releases/download/v2.10.11/nats-server-v2.10.11-linux-amd64.tar.gz \
    | tar xz && \
    mv nats-server-v2.10.11-linux-amd64/nats-server /usr/local/bin/ && \
    rm -rf nats-server-v2.10.11-linux-amd64

# Copiar app Go compilado
COPY --from=builder /app/app /usr/local/bin/app

# Copiar e dar permissão ao start.sh
COPY start.sh /usr/local/bin/start.sh
RUN chmod +x /usr/local/bin/start.sh

EXPOSE 8080 4222

# Starta via script
CMD ["/usr/local/bin/start.sh"]
