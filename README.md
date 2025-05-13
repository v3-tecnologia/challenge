<p align="center">
    <img src="./.github/logo.png" width="200px">
</p>

<h1 align="center" style="font-weight: bold;">Desafio Técnico da V3</h1>

## Dados para contato

LinkedIn := [https://www.linkedin.com/in/ricardoraposoo](https://www.linkedin.com/in/ricardoraposoo)
<br/>
Email := raposo.bomtempo@gmail.com

## 📝 Descrição

Este projeto tem como objetivo criar um sistema de registro e armazenamento de dados de telemetria, incluindo dados de giroscópio, GPS e fotos de usuários.

A aplicação foi desenvolvida em Go, utilizando o framework web [Fiber](https://gofiber.io/), e armazena os dados em um banco de dados PostgreSQL. Para gerar os métodos de acesso ao banco, utilizei o [sqlc](https://docs.sqlc.dev/), que permite manter a clareza do SQL com tipagem forte no Go e organização por repositórios. 

As migrações são gerenciadas com o [golang-migrate](https://github.com/golang-migrate/migrate), garantindo consistência no schema do banco de dados.

As imagens dos usuários são armazenadas em um bucket S3 da AWS, integrando com o [AWS Rekognition](https://aws.amazon.com/rekognition/) para análise de imagens, conforme exigido pelo desafio.

Acredito que todos os requisitos do desafio foram atendidos.

## 🛠 Tecnologias utilizadas

- [Go](https://golang.org/) – linguagem principal do projeto
- [Fiber](https://gofiber.io/) – framework web para Go
- [PostgreSQL](https://www.postgresql.org/) – banco de dados relacional
- [sqlc](https://docs.sqlc.dev/) – geração de código a partir de SQL
- [golang-migrate](https://github.com/golang-migrate/migrate) – controle de migrações do banco de dados
- [AWS S3](https://aws.amazon.com/s3/) – armazenamento de arquivos
- [AWS Rekognition](https://aws.amazon.com/rekognition/) – análise de imagens
- [Testify](https://github.com/stretchr/testify) – framework de testes
- [Docker](https://www.docker.com/) – containerização da aplicação

## Pré Requisitos

- Docker
- AWS S3 Bucket
- AWS IAM User

## 🚀 Executando o projeto

Como a configuração com a AWS possa ser um pouco exaustiva, eu gravei o vídeo abaixo demonstrando as funcionalidades de upload de imagem e detecção de rostos:


https://github.com/user-attachments/assets/03b5a114-d2b7-418b-98ad-528c814de55f


Os dois outros endpoints podem ser facilmente executados apenas com o banco de dados configurado.

### Configuração de ambiente

Crie um arquivo `.env` a partir do `.env.example` e preencha os campos de acordo com suas credenciais AWS e configurações do S3.
Algumas variáveis já estão pré preenchidas, para facilitar a execução do projeto.

### Migrações

Para controle das migrações, é necessário ter instalado a ferramenta [golang-migrate](https://github.com/golang-migrate/migrate).
Comando para executar as migrações:
```bash
migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path db/migrations up
```

### Execução do projeto

Para executar o projeto, é necessário ter instalado o [Docker](https://docs.docker.com/get-docker/).
Tendo feito as configurações prévias, para executar o projeto, basta executar o comando:
```bash
docker-compose up
```


##  📝 Executando os testes

Para execução dos testes, é necessário ter instalado as dependências do projeto.
Para isso, basta executar o comando:
```bash
go mod download
go mod tidy
```

Para executar os testes, basta executar o comando:
```bash
go test ./...
# ou 
make test
```
