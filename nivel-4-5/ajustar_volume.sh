#!/bin/bash
# ----------------------------------------
# Script para envio paralelo de comandos de ajuste de volume a dispositivos embarcados
# Com validações e tratamento de falhas de rede
# Autor: Guilherme Vasconcelos
# ----------------------------------------

mostrar_ajuda() {
  echo "Uso: $0 -f <arquivo.csv> -u <url_endpoint>"
  echo
  echo "  -f   Caminho para o arquivo CSV com placas e volumes"
  echo "  -u   URL do endpoint para envio dos comandos"
  echo "  -h   Exibe esta ajuda"
  exit 1
}

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

LOG_ARQ="log_envio_$(date +%Y%m%d_%H%M%S).log"
ARQ_OK=$(mktemp)
ARQ_FAIL=$(mktemp)

echo "📘 Iniciando envio de comandos para o endpoint: $ENDPOINT_URL" | tee -a "$LOG_ARQ"
echo "📂 Lendo arquivo: $ARQUIVO_CSV" | tee -a "$LOG_ARQ"
echo "--------------------------------------------" | tee -a "$LOG_ARQ"

processar_linha() {
  local linha="$1"
  local idx="$2"
  local endpoint="$3"
  local log_file="$4"

  IFS=',' read -r placa volume_alerta <<< "$linha"
  placa=$(echo "$placa" | xargs | tr '[:lower:]' '[:upper:]')
  volume_alerta=$(echo "$volume_alerta" | xargs)

  echo "[$(date '+%Y-%m-%d %H:%M:%S')] 🚚 Processando veículo $idx: placa=$placa, volume=$volume_alerta" >> "$log_file"

  if [[ ! "$placa" =~ ^[A-Z]{3}[0-9]{4}$ ]]; then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] ⚠️  Placa inválida: $placa. Pulando..." >> "$log_file"
    echo "1" >> "$ARQ_FAIL"
    return
  fi

  if [[ ! "$volume_alerta" =~ ^[0-9]+$ ]] || [ "$volume_alerta" -lt 0 ] || [ "$volume_alerta" -gt 101 ]; then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] ⚠️  Volume inválido ($volume_alerta) para placa $placa. Pulando..." >> "$log_file"
    echo "1" >> "$ARQ_FAIL"
    return
  fi

  resposta=$(curl --fail -s -m 10 --connect-timeout 5 -o /dev/null -w "%{http_code}" -X POST "$endpoint" \
    -H "Content-Type: application/json" \
    -d "{\"placa\": \"$placa\", \"volume_alerta\": $volume_alerta}")

  if [ "$?" -ne 0 ]; then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] ❌ Erro de conexão ao enviar para $placa" >> "$log_file"
    echo "1" >> "$ARQ_FAIL"
    return
  fi

  if [ "$resposta" -eq 200 ]; then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] ✅ Sucesso: comando enviado para $placa (volume $volume_alerta)" >> "$log_file"
    echo "1" >> "$ARQ_OK"
  else
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] ❌ Falha HTTP $resposta para $placa" >> "$log_file"
    echo "1" >> "$ARQ_FAIL"
  fi
}

export -f processar_linha
export ARQ_OK ARQ_FAIL

# Remove cabeçalho, indexa linhas e executa em paralelo com até 4 processos
tail -n +2 "$ARQUIVO_CSV" | nl -n ln | \
  xargs -L1 -P4 -I{} bash -c 'processar_linha "$(echo "{}" | cut -f2-)" "$(echo "{}" | cut -f1)" "'"$ENDPOINT_URL"'" "'"$LOG_ARQ"'"'

# Resultados
TOTAL=$(($(wc -l < "$ARQUIVO_CSV") - 1))
SUCESSO=$(wc -l < "$ARQ_OK")
FALHA=$(wc -l < "$ARQ_FAIL")

rm -f "$ARQ_OK" "$ARQ_FAIL"

echo "--------------------------------------------" | tee -a "$LOG_ARQ"
echo "📊 Execução finalizada." | tee -a "$LOG_ARQ"
echo "Total de veículos processados: $TOTAL" | tee -a "$LOG_ARQ"
echo "Comandos enviados com sucesso: $SUCESSO" | tee -a "$LOG_ARQ"
echo "Falhas durante o envio:        $FALHA" | tee -a "$LOG_ARQ"

echo "📁 Logs salvos em: $LOG_ARQ"
