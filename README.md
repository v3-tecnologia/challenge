# Cloud Challenge V3

# Descrição

Projeto desenvolvido utilizando o framework **GIN** para a API REST.  
Os dados são armazenados em um banco PostgreSQL dockerizado.  

A aplicação usa o serviço **migrate** no `docker-compose.yml` para aplicar as migrações do banco automaticamente durante o start.  

Testes completos foram implementados utilizando:  
- `httptest` para handlers HTTP,  
- `sqlmock` para simular o banco de dados,  
- `testify` para asserções,  
- `validator/v10` para garantir a integridade e validação dos dados.  

As requisições são validadas pelo Gin, com timestamps enviados no formato UNIX int64 e convertidos para datetime no banco.  
Somente MAC addresses válidos são aceitos.  

As fotos são armazenadas em um bucket S3, com as URLs salvas no banco de dados. Novas fotos enviadas são comparadas com as já existentes utilizando AWS Rekognition para reconhecimento facial.  

O tratamento de erros é realizado via erros customizados que permitem diferenciar falhas na API, banco de dados e serviços externos, garantindo mensagens claras e consistentes para o cliente.

Implementação de logging estruturado com o pacote log/slog, registrando eventos em formato JSON no terminal, facilitando a análise de requisições e erros em produção ou durante o desenvolvimento.

# Rodando Projeto

## Configuração do arquivo de ambiente  
Crie seu próprio arquivo `.env` seguindo o exemplo `.example.env` .

## Iniciar containers  
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

### 3. Photo (Deve ser um base64 valid de uma imagem com rosto, o abaixo é somente um exemplo)
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
