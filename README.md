<p align="center">
    <img src="./.github/logo.png" width="200px">
</p>

<h1 align="center" style="font-weight: bold;">Desafio Técnico da V3</h1>

## Tecnologias utilizadas

- GoLang

  - A aplicação foi construida em GoLang

- GORM

  - ORM utilizado para facilitar nas execuções de consultas no banco de dados

- Gin Framework

  - Framework foi utilizado para modelar a estrutura de RestAPI

- AWS Rekognition

  - Serviço utilizado para o reconhecimento facial de fotos recebidas na aplicação

- Swagger

  - Biblioteca utilizada para a documentação das rotas da aplicação, facilitando os testes em ambiente de desenvolvimento por ter uma integração pronta com o Gin Framework

- Migrate

  - Utilizado a biblioteca `migrate` para administrar as migrations do projeto.

##

## Como rodar localmente

Primeiro instale as dependências do projeto.

```shell
go mod install
```

Agora é necessário criar o arquivo `.env` no diretório raiz do projeto, copie o exemplo `.env.example` e preencha com as credenciais da AWS para funcionar o AWS Rekognition.

Você pode rodar essa aplicação localmente utilizando o docker.

```sh
docker compose up -d
```

Após subir a aplicação, execute as migrations para sincronizar com o banco de dados.

```sh
make migrate-up
```

OU

```shell
migrate -path ./internal/infra/database/migrations -database "postgres://dev:dev@localhost:5432/challenge_db?sslmode=disable" --verbose up
```

Caso não tenha o `go-migrate` instalado, basta seguir as instruções desse link:

```sh
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
```

## Como testar

Para rodar os testes da aplicação execute o seguinte comando.

```sh
go test ./...
```

E para fazer o teste da aplicação rodando localmente você pode estar acessando o seguinte endereço.

```shell
http://localhost:8080/swagger/index.html
```
