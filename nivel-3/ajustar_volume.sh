#!/bin/bash
# Este script foi gerado com a ajuda de IA (ChatGPT). Mas eu revisei, testei e ajustei.
# ----------------------------------------
# Script para envio paralelo de comandos de ajuste de volume a dispositivos embarcados
# Autor: Guilherme Vasconcelos
# Uso: ./ajustar_volume.sh -f arquivo.csv -u https://meu-endpoint.com/api/dispositivos/configurar
# ----------------------------------------

mostrar_ajuda() {
  echo "Uso: $0 -f <arquivo.csv> -u <url_endpoint>"
  echo
  echo "  -f   Caminho para o arquivo CSV com placas e volumes"
  echo "  -u   URL do endpoint para envio dos comandos"
  echo "  -h   Exibe esta ajuda"
  exit 1
}

# ----------------------------------------
# Processa os parâmetros passados para o script via linha de comando.
# Aceita as seguintes opções:
#   -f <arquivo.csv>   : Caminho para o arquivo CSV com placas e volumes
#   -u <url_endpoint>  : URL do endpoint para envio dos comandos
#   -h                 : Exibe a mensagem de ajuda
# ----------------------------------------
while getopts ":f:u:h" opt; do
  case ${opt} in
    f ) ARQUIVO_CSV="$OPTARG" ;;
    u ) ENDPOINT_URL="$OPTARG" ;;
    h ) mostrar_ajuda ;;
    \? ) echo "Parâmetro inválido: -$OPTARG" >&2; mostrar_ajuda ;;
    : ) echo "Opção -$OPTARG requer um argumento." >&2; mostrar_ajuda ;;
  esac
done

if [ -z "$ARQUIVO_CSV" ] || [ -z "$ENDPOINT_URL" ]; then
  echo "❌ Erro: arquivo CSV e URL do endpoint são obrigatórios."
  mostrar_ajuda
fi

if [ ! -f "$ARQUIVO_CSV" ]; then
  echo "❌ Erro: Arquivo '$ARQUIVO_CSV' não encontrado."
  exit 1
fi

echo "📘 Iniciando envio de comandos para o endpoint: $ENDPOINT_URL"
echo "📂 Lendo arquivo: $ARQUIVO_CSV"
echo "--------------------------------------------"

# Arquivos temporários para contagem
ARQ_OK=$(mktemp)
ARQ_FAIL=$(mktemp)

processar_linha() {
  local linha="$1"
  local idx="$2"
  local endpoint="$3"

  IFS=',' read -r placa volume_alerta <<< "$linha"
  placa=$(echo "$placa" | xargs)
  volume_alerta=$(echo "$volume_alerta" | xargs)

  echo "🚚 Processando veículo $idx: placa=$placa, volume=$volume_alerta ..."

  if [[ ! "$volume_alerta" =~ ^[0-9]+$ ]] || [ "$volume_alerta" -lt 0 ] || [ "$volume_alerta" -gt 100 ]; then
    echo "⚠️  Volume inválido ($volume_alerta) para a placa $placa. Pulando..."
    echo "1" >> "$ARQ_FAIL"
    echo "--------------------------------------------"
    return
  fi

  resposta=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$endpoint" \
    -H "Content-Type: application/json" \
    -d "{\"placa\": \"$placa\", \"volume_alerta\": $volume_alerta}")

  if [ "$resposta" -eq 200 ]; then
    echo "✅ Sucesso: Comando enviado para $placa (volume $volume_alerta)"
    echo "1" >> "$ARQ_OK"
  else
    echo "❌ Falha: HTTP $resposta ao enviar comando para $placa"
    echo "1" >> "$ARQ_FAIL"
  fi
  echo "--------------------------------------------"
}

export -f processar_linha
export ARQ_OK ARQ_FAIL

# Remove cabeçalho e executa linhas em paralelo com xargs
tail -n +2 "$ARQUIVO_CSV" | nl -n ln | \
  xargs -L1 -P4 -I{} bash -c 'processar_linha "$(echo "{}" | cut -f2-)" "$(echo "{}" | cut -f1)" "'"$ENDPOINT_URL"'"'

# Contagem final
TOTAL=$(($(wc -l < "$ARQUIVO_CSV") - 1))
SUCESSO=$(wc -l < "$ARQ_OK")
FALHA=$(wc -l < "$ARQ_FAIL")

rm -f "$ARQ_OK" "$ARQ_FAIL"

echo "📊 Execução finalizada."
echo "Total de veículos processados: $TOTAL"
echo "Comandos enviados com sucesso: $SUCESSO"
echo "Falhas durante o envio:        $FALHA"
