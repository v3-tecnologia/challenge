# 🛰️ V3 Challenge API

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.24-00ADD8?logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-ready-2496ED?logo=docker&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-Rekognition-FF9900?logo=amazon-aws&logoColor=white)

## 📋 Visão Geral

V3 Challenge API é uma API robusta desenvolvida para receber, validar e processar dados de telemetria de dispositivos, incluindo informações de giroscópio, GPS e imagens. A API utiliza tecnologias modernas como Golang, Docker e integração com AWS Rekognition para reconhecimento facial avançado.

## 🚀 Funcionalidades

- ✅ Recebimento e validação de dados de telemetria (giroscópio, GPS, fotos)
- 📦 Armazenamento consistente em banco de dados
- 🧠 Reconhecimento inteligente de imagens com AWS Rekognition
- 🔍 Verificação de similaridade entre fotos enviadas
- 🐳 Containerização completa com Docker
- 🧪 Cobertura completa de testes unitários

## 📚 Endpoints da API

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `POST` | `/telemetry/gyroscope` | Recebe e processa dados do giroscópio |
| `POST` | `/telemetry/gps` | Recebe e processa dados de GPS |
| `POST` | `/telemetry/photo` | Recebe, processa e analisa fotos |

## 📐 Estrutura dos Dados

### Giroscópio
```json
{
  "mac_address": "string",
  "timestamp": "Time",
  "x": "number",
  "y": "number",
  "z": "number"
}
```

### GPS
```json
{
  "mac_address": "string",
  "timestamp": "Time",
  "latitude": "number",
  "longitude": "number",
}
```

### Foto
```json
{
  "mac_address": "string",
  "timestamp": "Time",
  "file_url": "string",
  "is_match": "bool"
}
```

## ❌ Erros possíveis
### 500
![image](https://github.com/user-attachments/assets/1e40bde4-5507-49ab-a2b0-232d1d672031)

### 400
![image](https://github.com/user-attachments/assets/6214cb9c-9796-4fac-bd2b-2104997712ec)

## 🖥️ Tecnologias Utilizadas

- **Backend**: Golang
- **Banco de Dados**: Postgres
- **Containerização**: Docker, Docker Compose
- **Testes**: Go testing, Go Sql Mock
- **Reconhecimento de Imagem**: AWS Rekognition

## 🛠️ Configuração e Instalação

### Para rodar o projeto localmente utilize e consulte o arquivo ```Makefile```

- Para rodar e buildar:
```make
  make run
```

- Para buildar
```make
  make build
```

- Para rodar os testes:
```make
  make test
```

- Para subir o container com a aplicação e o banco
```docker
  docker compose up --build
```

### Pré-requisitos

- Go (v1.24 ou superior)
- Docker e Docker Compose
- Conta AWS (para Rekognition)

Desenvolvido com ❤️ para a V3.
