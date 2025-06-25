#!/bin/sh

# Iniciar NATS em background
echo "Iniciando NATS Server..."
nats-server --port 4222 --http_port 8222 &

# Aguardar NATS inicializar
sleep 3

# Verificar se NATS está rodando
echo "Verificando se NATS está ativo..."
if ! nc -z localhost 4222; then
    echo "Erro: NATS não iniciou corretamente"
    exit 1
fi

echo "NATS Server iniciado com sucesso!"

# Configurar variáveis de ambiente
export NATS_URL="nats://localhost:4222"
export PORT="${PORT:-8080}"

# Iniciar a API
echo "Iniciando API Golang..."
exec ./api