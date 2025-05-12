# Aplicação de Telemetria para Veículos

<div align="center">
  <img src="https://img.shields.io/badge/Go-1.24.2-blue" alt="Go Version" />
  <img src="https://img.shields.io/badge/License-MIT-green" alt="License" />
</div>

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Funcionalidades](#funcionalidades)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Requisitos](#requisitos)
- [Configuração](#configuração)
- [Como Executar](#como-executar)
    - [Docker (Recomendado)](#usando-docker-recomendado)
    - [Localmente](#execução-local)
- [API](#api)
    - [Endpoints](#endpoints-da-api)
    - [Exemplos de Requisições](#exemplos-de-requisições)
- [Testes](#testes)
- [Recursos Avançados](#recursos-avançados)
- [Contribuição](#contribuição)
- [Licença](#licença)
- [Contato](#contato)

## 🌟 Visão Geral

_Este projeto é um MVP (Minimum Viable Product) para um sistema de telemetria veicular que utiliza celulares Android como dispositivos de coleta de dados. O back-end é responsável por receber, processar e armazenar os dados de telemetria enviados pelos dispositivos móveis, tornando-os disponíveis por meio de uma API REST._

## ✨ Funcionalidades

A aplicação coleta e armazena os seguintes tipos de dados:
1. **Dados de Giroscópio** - Valores de rotação nos eixos x, y e z, com timestamp
2. **Dados de GPS** - Latitude e longitude, com timestamp
3. **Fotos** - Imagens capturadas pela câmera do dispositivo, com timestamp
4. **Identificação do Dispositivo** - Endereço MAC ou outro identificador único

## 📁 Estrutura do Projeto

O projeto segue uma arquitetura limpa (Clean Architecture), organizada da seguinte forma:
- : Contém a lógica principal da aplicação
    - : Entidades e regras de negócio `/domain`
    - `/application`: Casos de uso da aplicação
    - : Controladores e apresentadores `/interface`
    - `/infrastructure`: Implementações concretas (banco de dados, serviços externos)
    - `/errors`: Tratamento de erros customizados
    - : Configuração das rotas da API `/router`

`/internal`
- : Arquivos de configuração `/config`
- `/pkg`: Bibliotecas e utilitários compartilhados
- `/migrations`: Scripts de migração do banco de dados
  o

## 📋 Requisitos

Go 1.23 ou superior
- Docker e Docker Compose
- AWS CLI configurada (para funcionalidades de reconhecimento de imagens)


## ⚙️ Configuração

git clone [URL_DO_REPOSITORIO]

## 🚀 Como Executar

### Usando Docker (Recomendado)
1. Certifique-se de ter o Docker e o Docker Compose instalados em sua máquina.
2. Execute o seguinte comando para construir e iniciar os containers:
``` bash
   docker-compose up -d
```
1. A API estará disponível em `http://localhost:8080`

### Execução Local
1. Certifique-se de ter Go 1.24.2 ou superior instalado.
2. Instale as dependências:
``` bash
   go mod tidy
```
1. Configure as variáveis de ambiente (consulte o arquivo `.env.example`)
2. Execute a aplicação:
``` bash
   go run main.go
```

## 🔌 API

### Endpoints da API

A API disponibiliza os seguintes endpoints para receber dados de telemetria:
- `POST /telemetry/gyroscope` - Recebe dados do giroscópio
- `POST /telemetry/gps` - Recebe dados do GPS
- `POST /telemetry/photo` - Recebe fotos capturadas pela câmera

### Exemplos de Requisições
**Giroscópio:**
``` json
POST /telemetry/gyroscope
{
  "device_id": "string",
  "created_at": "2023-01-01T12:00:00Z",
  "x": 0.0,
  "y": 0.0,
  "z": 0.0
}
```
**GPS:**
``` json
POST /telemetry/gps
{
  "device_id": "string",
  "created_at": "2023-01-01T12:00:00Z",
  "latitude": 0.0,
  "longitude": 0.0
}
```
**Foto:**
``` 
POST /telemetry/photo
Content-Type: multipart/form-data

- device_id: string
- created_at: string (formato ISO 8601)
- photo: arquivo de imagem
```

