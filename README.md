# API de Telemetria

Uma API minimalista para coleta de dados de telemetria de dispositivos, incluindo informações de giroscópio, GPS e fotos. Apresenta uma arquitetura orientada por eventos usando NATS JetStream, com um sistema de autenticação seguro usando JWT, tokens de atualização e controle de acesso baseado em funções.

## Fluxo da Arquitetura

O sistema é projetado com uma arquitetura assíncrona e desacoplada para garantir capacidade de processamento e resistência. A API atua como um gateway leve, publicando dados recebidos em uma fila NATS. Um serviço Worker separado processa os dados da fila e os persiste no banco de dados.

```mermaid
%%{init: {"theme": "dark", "themeVariables": {"primaryColor": "#9333ea", "primaryTextColor": "#e5e7eb", "primaryBorderColor": "#7c3aed", "lineColor": "#a855f7", "secondaryColor": "#581c87", "tertiaryColor": "#3730a3", "background": "#1f2937", "mainBkg": "#374151", "secondBkg": "#4b5563", "tertiaryBkg": "#6b7280"}}}%%
sequenceDiagram
    participant Cliente as Cliente
    participant Gateway as Gateway API
    participant NATS as NATS JetStream
    participant Worker as Worker
    participant DB as Banco de Dados

    Cliente->>Gateway: POST /v1/auth/register (dados do usuário)
    Gateway->>DB: Hash da senha e criar usuário
    DB-->>Gateway: Usuário criado
    Gateway-->>Cliente: 201 Created (detalhes do usuário)

    Cliente->>Gateway: POST /v1/auth/login (credenciais)
    Gateway->>DB: Encontrar usuário e verificar senha
    DB-->>Gateway: Usuário autenticado
    Gateway-->>Gateway: Gerar Tokens JWT de Acesso e Atualização
    Gateway->>DB: Armazenar Token de Atualização
    Gateway-->>Cliente: 200 OK (tokens, funções)

    Cliente->>Gateway: POST /v1/telemetry/gps (dados)<br/>Authorization: Bearer access_token
    Gateway->>Gateway: Validar JWT e verificar função 'admin'
    Gateway->>NATS: Publicar dados GPS no tópico "telemetry.gps"
    Gateway-->>Cliente: 202 Accepted

    NATS->>Worker: Enviar mensagem GPS
    Worker->>DB: Salvar dados de telemetria
    DB-->>Worker: Dados salvos
    Worker->>NATS: Confirmar mensagem
```

## Começando

Siga estes passos para colocar a API de Telemetria em funcionamento na sua máquina local.

### Pré-requisitos

- Docker
- Docker Compose

### Executando a Aplicação

1.  **Clone o repositório:**

    ```bash
    git clone <url-do-repositório>
    cd telemetry-api
    ```

2.  **Configure as variáveis de ambiente (opcional):**

    Para desenvolvimento local, você pode criar um arquivo `.env` na raiz do projeto com as seguintes configurações:

    ```env
    DB_HOST=localhost
    POSTGRES_USER=telemetry_user
    POSTGRES_PASSWORD=telemetry_P0stgr3s_2025!
    POSTGRES_DB=telemetry_db
    JWT_SECRET=404931fa-aa9c-48f9-b576-deab7b54d42b
    ```

    > **Nota:** O Docker Compose já inclui essas configurações por padrão. O arquivo `.env` é útil apenas se você quiser executar a aplicação fora do Docker.

3.  **Inicie os serviços:**
    Use o Docker Compose para construir e executar os serviços da API, worker e banco de dados.
    ```bash
    docker-compose up --build
    ```
    A API estará disponível em `http://localhost:8080`.

## Endpoints da API

Todos os endpoints são prefixados com `/v1`.

### Autenticação

#### Registrar um novo usuário

```bash
curl -X POST http://localhost:8080/v1/auth/register \
-H "Content-Type: application/json" \
-d '{
  "username": "admin",
  "email": "admin@example.com",
  "password": "password123"
}'
```

#### Login

```bash
curl -X POST http://localhost:8080/v1/auth/login \
-H "Content-Type: application/json" \
-d '{
  "email": "admin@example.com",
  "password": "password123"
}'
```

_Retorna um `access_token` e um `refresh_token`._

#### Token de Atualização

```bash
curl -X POST http://localhost:8080/v1/auth/refresh \
-H "Content-Type: application/json" \
-d '{
  "refresh_token": "seu_token_de_atualizacao_aqui"
}'
```

### Telemetria (Somente Admin)

Estes endpoints requerem um `access_token` válido no cabeçalho `Authorization` e a função `admin`.

#### Criar dados (POST)

Estes endpoints aceitam os dados e os colocam em fila para processamento assíncrono.

#### Enviar dados do giroscópio

```bash
curl -X POST http://localhost:8080/v1/telemetry/gyroscope \
-H "Content-Type: application/json" \
-H "Authorization: Bearer SEU_TOKEN_DE_ACESSO_AQUI" \
-d '{
  "device_id": "AA:BB:CC:DD:EE:FF",
  "x": 1.23,
  "y": 4.56,
  "z": 7.89,
  "timestamp": "2024-06-24T10:30:00Z"
}'
```

_Retorna `202 Accepted` para indicar que a solicitação foi colocada em fila._

#### Enviar dados GPS

```bash
curl -X POST http://localhost:8080/v1/telemetry/gps \
-H "Content-Type: application/json" \
-H "Authorization: Bearer SEU_TOKEN_DE_ACESSO_AQUI" \
-d '{
  "device_id": "AA:BB:CC:DD:EE:FF",
  "latitude": -23.5505,
  "longitude": -46.6333,
  "timestamp": "2024-06-24T10:30:00Z"
}'
```

_Retorna `202 Accepted` para indicar que a solicitação foi colocada em fila._

#### Enviar Dados de Foto

```bash
# O corpo deve conter a string da imagem codificada em base64
curl -X POST http://localhost:8080/v1/telemetry/photo \
-H "Content-Type: application/json" \
-H "Authorization: Bearer SEU_TOKEN_DE_ACESSO_AQUI" \
-d '{
  "device_id": "AA:BB:CC:DD:EE:FF",
  "timestamp": "2024-06-24T10:30:00Z",
  "photo": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII="
}'

# Pode usar este para testar:
# iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=
```

_Retorna `202 Accepted` para indicar que a solicitação foi colocada em fila._

#### Consultar dados (GET)

Estes endpoints permitem consultar os dados de telemetria armazenados com suporte a paginação e filtros.

#### Consultar dados do giroscópio

```bash
# Listar todos os dados de giroscópio (paginado)
curl -X GET "http://localhost:8080/v1/telemetry/gyroscope?page=1&limit=10" \
-H "Authorization: Bearer SEU_TOKEN_DE_ACESSO_AQUI"

# Filtrar por dispositivo específico
curl -X GET "http://localhost:8080/v1/telemetry/gyroscope?device_id=AA:BB:CC:DD:EE:FF&page=1&limit=5" \
-H "Authorization: Bearer SEU_TOKEN_DE_ACESSO_AQUI"
```

#### Consultar dados GPS

```bash
# Listar todos os dados GPS (paginado)
curl -X GET "http://localhost:8080/v1/telemetry/gps?page=1&limit=10" \
-H "Authorization: Bearer SEU_TOKEN_DE_ACESSO_AQUI"

# Filtrar por dispositivo específico
curl -X GET "http://localhost:8080/v1/telemetry/gps?device_id=AA:BB:CC:DD:EE:FF" \
-H "Authorization: Bearer SEU_TOKEN_DE_ACESSO_AQUI"
```

#### Consultar dados de fotos

```bash
# Listar todos os dados de fotos (paginado)
curl -X GET "http://localhost:8080/v1/telemetry/photo?page=1&limit=10" \
-H "Authorization: Bearer SEU_TOKEN_DE_ACESSO_AQUI"

# Filtrar por dispositivo específico
curl -X GET "http://localhost:8080/v1/telemetry/photo?device_id=AA:BB:CC:DD:EE:FF" \
-H "Authorization: Bearer SEU_TOKEN_DE_ACESSO_AQUI"
```

#### Listar dispositivos únicos

```bash
# Listar todos os dispositivos com estatísticas de uso
curl -X GET "http://localhost:8080/v1/telemetry/devices?page=1&limit=20" \
-H "Authorization: Bearer SEU_TOKEN_DE_ACESSO_AQUI"
```

**Parâmetros de consulta disponíveis:**

- `page` (opcional): Número da página (padrão: 1)
- `limit` (opcional): Itens por página (padrão: 10, máximo: 100)
- `device_id` (opcional): Filtrar por ID do dispositivo específico

**Estrutura de resposta:**

```json
{
  "data": [
    {
      "id": "uuid-do-registro",
      "device_id": "AA:BB:CC:DD:EE:FF",
      "timestamp": "2024-06-24T10:30:00Z",
      "created_at": "2024-06-24T10:30:05Z"
      // ... campos específicos do tipo de dados
    }
  ],
  "pagination": {
    "current_page": 1,
    "per_page": 10,
    "total": 150,
    "total_pages": 15
  }
}
```

**Resposta do endpoint /devices:**

```json
{
  "data": [
    {
      "device_id": "AA:BB:CC:DD:EE:FF",
      "last_seen": "2024-06-24T10:30:00Z",
      "gyroscope_count": 45,
      "gps_count": 32,
      "photo_count": 12,
      "total_data_points": 89
    }
  ],
  "pagination": {
    "current_page": 1,
    "per_page": 10,
    "total": 5,
    "total_pages": 1
  }
}
```

## Monitoramento da Fila NATS

O NATS JetStream fornece endpoints HTTP para monitoramento em tempo real das filas e do processamento de mensagens.

### Status Geral do Servidor

```bash
# Informações gerais do servidor NATS
curl http://localhost:8222/varz
```

### Informações do JetStream

```bash
# Status das streams e consumidores
curl http://localhost:8222/jsz

# Informações detalhadas das streams
curl http://localhost:8222/jsz?streams=true

# Informações detalhadas dos consumidores
curl http://localhost:8222/jsz?consumers=true
```

### Conexões e Subscrições

```bash
# Conexões ativas
curl http://localhost:8222/connz

# Subscrições ativas
curl http://localhost:8222/subsz
```

## Recursos Principais

- **Arquitetura Orientada por Eventos**: Usa NATS JetStream para enfileiramento de mensagens assíncrono e confiável entre a API e um worker em segundo plano.
- **Autenticação JWT**: Tokens de acesso de curta duração (60 minutos) e tokens de atualização de longa duração (30 dias).
- **Controle de Acesso Baseado em Funções (RBAC)**: Endpoints de telemetria são restritos a usuários com a função `admin`.
- **Hash de Senhas**: Senhas são hasheadas com segurança usando bcrypt.
- **Limitação de Taxa**: Protege contra ataques de força bruta e negação de serviço.
- **Validação de Entrada**: Todos os corpos de solicitação são validados no gateway da API.
