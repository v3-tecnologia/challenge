#!/bin/bash
# Este script foi gerado com a ajuda de IA (deepseek). Mas eu revisei, testei e ajustei.
# ----------------------------------------
# Script para envio de comandos de ajuste de volume a dispositivos embarcados
# Autor: Guilherme Vasconcelos
# Uso: ./ajustar_volume.sh -p PLACA -v VOLUME -u https://meu-endpoint.com/api/dispositivos/configurar
# ----------------------------------------

mostrar_ajuda() {
  echo "Uso: $0 -p <placa> -v <volume> -u <url_endpoint>"
  echo
  echo "  -p   Placa do veículo"
  echo "  -v   Volume a ser definido (0-100)"
  echo "  -u   URL do endpoint para envio dos comandos"
  echo "  -h   Exibe esta ajuda"
  exit 1
}

# ----------------------------------------
# Processa os parâmetros da linha de comando
# ----------------------------------------
while getopts ":p:v:u:h" opt; do
  case ${opt} in
    p ) PLACA="$OPTARG" ;;
    v ) VOLUME="$OPTARG" ;;
    u ) ENDPOINT_URL="$OPTARG" ;;
    h ) mostrar_ajuda ;;
    \? ) echo "Parâmetro inválido: -$OPTARG" >&2; mostrar_ajuda ;;
    : ) echo "Opção -$OPTARG requer um argumento." >&2; mostrar_ajuda ;;
  esac
done

# Trima espaços em branco dos parâmetros
PLACA=$(echo "$PLACA" | xargs)
VOLUME=$(echo "$VOLUME" | xargs)
ENDPOINT_URL=$(echo "$ENDPOINT_URL" | xargs)

# Validação dos parâmetros obrigatórios
if [ -z "$PLACA" ] || [ -z "$VOLUME" ] || [ -z "$ENDPOINT_URL" ]; then
  echo "❌ Erro: placa, volume e URL do endpoint são obrigatórios."
  mostrar_ajuda
fi

echo "📘 Iniciando envio de comando para o endpoint: $ENDPOINT_URL"
echo "🚚 Processando veículo: placa=$PLACA, volume=$VOLUME ..."
echo "--------------------------------------------"

# Validação do volume
if [[ ! "$VOLUME" =~ ^[0-9]+$ ]] || [ "$VOLUME" -lt 0 ] || [ "$VOLUME" -gt 100 ]; then
  echo "⚠️  Volume inválido ($VOLUME) para a placa $PLACA. Pulando..."
  FALHA=1
else
  resposta=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$ENDPOINT_URL" \
    -H "Content-Type: application/json" \
    -d "{\"placa\": \"$PLACA\", \"volume_alerta\": $VOLUME}")

  if [ "$resposta" -eq 200 ]; then
    echo "✅ Sucesso: Comando enviado para $PLACA (volume $VOLUME)"
    SUCESSO=1
  else
    echo "❌ Falha: HTTP $resposta ao enviar comando para $PLACA"
    FALHA=1
  fi
fi

echo "--------------------------------------------"
echo "📊 Execução finalizada."

exit 0