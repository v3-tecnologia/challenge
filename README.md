# Desafio **CLOUD** da **V3 Tecnologia**.

**Objetivo:** Desenvolver um sistema (API) em Go que receba informações enviadas por um aplicativo (giroscópio, gps e foto) em um celular.

### Execução da **aplicação**
Para executar a aplicação execute o comando:
```
git clone https://github.com/IgorLopes88/desafio-v3.git
cd desafio-v3
docker compose up -d
```

O resultado deverá ser esse:

```
 ✔ Container mysql-v3  Started
 ✔ Container pma-v3    Started 
```

Rode a API

```
cd cmd/server
go run main.go
```
Acessar API: http://localhost:8000

Acessar PhpMyAdmin: http://localhost:8080

Acessar Swagger: http://localhost:8000/docs/ `(incompleto)`



# Etapas Concluidas

## Nível 1 - API e Validação

* ✔ `POST /telemetry/gyroscope` - Dados do giroscópio;
* ✔ `POST /telemetry/gps` - Dados do GPS;
* ✔ `POST /telemetry/photo` - Dados da Foto; `incompleto`

Foi incluído o cadastro de usuário/gerar token de acesso para essas rotas. Cadastro de fotos foi pensado duas formas: **json**, assim a imagem capturada precisa ser enviada como byte, podendo ser salva no banco de dados (não recomendado) ou convertida como arquivo; OU *multipart/form-data* onde a imagem é enviada de forma convencional (através de formulário).

## Nível 2 - Persistência e Testes

* ✔ Salve cada uma das informações em um banco de dados a sua escolha.
* ✔ Salve estes dados de forma identificável e consistente.
* ✔ Crie testes unitários para cada arquivo da aplicação.
* **X** Implemente testes de integração para validar o fluxo completo.

O banco de dados escolhido foi mysql/mariadb. Como foi adicionado o ID do usuário + MAC nas informações coletadas é possivel localizar todas as informações (relacionamento). Foi criados os arquivos de testes, exceto para a entidade de Photo.

## Nível 3 - Containerização e CI/CD

* ✔ Crie um _container_ em _Docker_ que contenha a sua aplicação e o banco de dados.
* **X** Configure um pipeline de CI/CD (pode usar GitHub Actions, GitLab CI, ou similar).
* **X** Implemente testes automatizados no pipeline.
* **X** Configure o processo de build e deploy da aplicação.

## Nível 4 - Processamento de Imagens

* **X** Implemente a integração com AWS Rekognition para análise de fotos.
* **X** Compare cada nova foto com as fotos anteriores enviadas.
* **X** Retorne um atributo indicando se a foto foi reconhecida.
* **X** Implemente um sistema de cache para otimizar as consultas ao Rekognition.
* **X** Adicione logs detalhados do processo de análise.

## Nível 5 - Arquitetura com Filas

* **X** Implemente um sistema de mensageria (RabbitMQ, Apache Kafka, AWS SQS, NATS (usamos este)):
  * **X** Crie filas separadas para cada tipo de telemetria
  * **X** Implemente o padrão Producer-Consumer
  * **X** Configure retry policies para mensagens com erro
  * **X** Implemente dead letter queues para mensagens problemáticas

## Nível 6 - Monitoramento e Documentação

* **X** Métricas de performance (Prometheus/Grafana)
* **X** Logs estruturados (ELK Stack ou similar)
* **X** Tracing distribuído
* **X** Alertas para condições anormais
* **X** Diagramas da arquitetura
* **X** Documentação técnica detalhada
* **X** Guia de operação e manutenção
* **X** Procedimentos de backup e recuperação
* **X** Guia de troubleshooting
* **✔** Documentação de APIs usando OpenAPI/Swagger
* **X** Guia de segurança e boas práticas
* **X** Procedimentos de escalabilidade e resiliência
* **X** Documentação de configuração e variáveis de ambiente
* **X** Guia de contribuição e desenvolvimento

Foi realizado a documentação da API com o Swagger. Basta acessar http://localhost:8000/docs/ ainda falta realizar alguns ajustes (inclusive validação dos campos). Mas já esta quase toda funcional. Como foi utilizado um *Web Router* para essa API. Tanto o Otel quanto RateLimiter podem sem fácilmente incluidos como Middleware. Tenho dois repositórios com essas implementações.

## Nível 7 - Segurança e Governança

* **✔**  Autenticação e autorização para todos os endpoints
* **X** Criptografia de dados sensíveis
* **X** Rate limiting para proteger a API
* **X** Validação de schema para todas as mensagens
* **X** Auditoria de todas as operações no sistema

Foi tutilizado o JWT para gerar a autenticação das rotas. Basta utilizar o arquivo `./test/user.http` criar um usuário, em seguida gerar um token. Esse token deverá se incluído nas requisições (tanto nos demais arquivos `gps.http` e `gyroscope.http`). O envio de fotos pode ser testado via Postman ou cURL

```
curl --request POST \
  --url http://localhost:8000/telemetry/photo \
  --header 'Authorization: Bearer <aqui seu token>' \
  --header 'Content-Type: multipart/form-data' \
  --form user=78aed56f-a85b-41a3-afa4-f16bf718a643 \
  --form mac_address=00:1A:2B:3C:4D:5E \
  --form 'timestamp=2025-06-30 10:00:10' \
  --form 'image=@C:\image.jpg'
```