# Cloud Challenge V3

# Descrição

Projeto realizado utilizando o framework GIN para a API rest.\
Dados salvos em um database postgres dockerizado.\
Testes unitários da api com httptest e do database com sqlmock

Checagem de requisição com gin. Timestamps são enviados como UNIX int64 e convertidos ao salvar no database, somente MACs validos podem ser enviados.\
Fotos são arquivadas como base64


# Rodando Projeto

## Configure o arquivo de ambiente:
Pode utilizar .example.env removendo .example do nome do arquivo. Ou crie seu arquivo .env seguindo o exemplo.

## Rode o container
```bash
docker compose up --build
```

## Utilizando a API
É possível importar ao postman a collection criada para esse projeto. [Challenge_Collection.postman_collection.json](Challenge_Collection.postman_collection.json)\
Seguem as requisições em CURL:
### 1. GPS
```bash
curl -X POST http://127.0.0.1:8080/telemetry/gps \
  -H "Content-Type: application/json" \
  -d '{
    "latitude": -23.5475,
    "longitude": 46.6361,
    "mac": "8C:16:45:8D:F3:7B",
    "timestamp": 1746110367207
  }'
```

### 2. Gyroscope
```bash
curl -X POST http://127.0.0.1:8080/telemetry/gyroscope \
  -H "Content-Type: application/json" \
  -d '{
    "x": 23.00,
    "y": 52.00,
    "z": 33.00,
    "mac": "8C:16:45:8D:F3:7B",
    "timestamp": 1746110367207
  }'

```

### 3. Photo
```bash
curl -X POST http://127.0.0.1:8080/telemetry/photo \
  -H "Content-Type: application/json" \
  -d '{
    "image_base_64": "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAABT0lEQVR4nO3ZsUoDQRBF0XcEJWJgEkVIsAcViMQFkpzDShvQhK4gYAvnQ5sHh51Znc7gUph/2Qzpd3Uz/d+dcBAAAAAAAAAAAAAPCxfMPu/O+zyPgL4Ir/w9YgHqwnV+Qj6kqY5H+aVVEY7VFEY7VVEY7VFEY7VVEY7VFEY7VVEY7VFEY7VVEY7VFEY7VXkofytVRGO1URjtVUQjtVURjtVUQjtVUQjtVURjtVURjtVURrtf8c36Dfsddz6F+2fs/kl9wP6Ph3sQAAAAAAAAAAAAB8HP4BgA+6aY5J1s8sAAAAAElFTkSuQmCC",
    "mac": "8C:16:45:8D:F3:7B",
    "timestamp": 1746110367207
  }'
```

## Rodando testes
Testes foram criados para os handlers das apis e os services

```bash
go test ./...
```
