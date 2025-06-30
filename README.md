# Desafio Firmware

Desafio: MVP para Coleta de Dados de Dispositivo Móvel

## 📌 Contexto do Negócio

Precisamos validar se smartphones Android podem substituir dispositivos IoT na coleta de dados de frotas. Seu objetivo é criar uma solução que:

- Rode continuamente em segundo plano, sem interface gráfica
- Capture giroscópio (x,y,z), GPS (lat/long) e fotos a cada 10 segundos
- Associe um identificador único do dispositivo
- Armazene e disponibilize os dados para integração
- Capture uma foto do motorista quando o sistema for acionado
- Permita a comunicação da Nuvem com o app para reconfigurações dinâmicas (***Diferencial***)

---

## ✅ Funcionalidades Implementadas

### 🎯 Coleta de Dados

- Captura automática de:
  - Localização (Latitude, Longitude)
  - Giroscópio (X, Y, Z)
  - Timestamp
  - Foto da câmera frontal
  - ID único do Android

- As coletas são salvas localmente no banco (Room) e também exportadas em `.bin` (Protobuf)

### 📸 Validação de Foto

- Verifica se a imagem tem rosto
- Verifica se a imagem não está toda preta/branca
- Se não tiver rosto, coleta é marcada como “FOTO SEM ROSTO” e enviada sem imagem

### 🌀 Serviço em Segundo Plano

- Serviço `ColetaService` roda como **foreground service** com notificação ativa
- Mantém o app vivo mesmo com a tela desligada
- Usa coroutines e `Job` para controle do loop

### ⏱️ Intervalo de Coleta

- Padrão: a cada 10 segundos
- Pode ser alterado via `Intent` (ação `ACTION_UPDATE_INTERVAL`)

### 🔄 Reenvio Automático

- Coletas com falha no envio são salvas com status `enviado = false`
- Ao iniciar o serviço, o app tenta reenviar todas as pendentes

### 🧯 Resiliência

- Localização inválida (null ou 0.0) é descartada e notificada via Toast
- Giroscópio ausente cancela a coleta e notifica
- Falhas inesperadas geram notificação para o usuário
- Logs de erro em `Logcat`

### 💾 Armazenamento e Exportação

- Cada coleta válida gera um `.bin` no diretório `/Download`
- Payload enviado via HTTP POST em `application/octet-stream`

---

## 🧪 Teste com Postman

1. Rode o backend local (Spring Boot na porta 8080)
2. Pegue um `.bin` gerado (em /Download)
3. No Postman:
  - `POST http://localhost:8080/api/binario`
  - Body: binary > escolha o `.bin`
  - Content-Type: `application/octet-stream`

---

## 📦 Protobuf

```proto
message ColetaMsg {
  string deviceId = 1;
  double latitude = 2;
  double longitude = 3;
  float gyroX = 4;
  float gyroY = 5;
  float gyroZ = 6;
  string status = 7;
  int64 timestamp = 8;
  bytes foto = 9;
}
