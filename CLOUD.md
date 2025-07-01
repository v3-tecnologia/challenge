# Desafio Cloud

## 💻 O Problema

Um dos nossos clientes ainda não consegue comprar o equipamento para colocar nos veículos de sua frota, mas ele quer muito utilizar a nossa solução.

Por isso, vamos fazer um MVP bastante simples para testar se, o celular do motorista poderia ser utilizado como o dispositivo de obtenção das informações.

> Parece fazer sentido certo? Ele possui vários mecanismos parecidos com o equipamento que oferecemos!

Sua missão ajudar na criação deste MVP para que possamos testar as frotas deste cliente.

Essa versão do produto será bastante simplificada. Queremos apenas criar as estruturas para obter algumas informações do seu dispositivo (Android) e armazená-la em um Banco de Dados.

Essas informações, depois de armazenadas devem estar disponíveis através de uma API para que este cliente integre com um Front-end já existente!

### Quais serão as informações que deverão ser coletadas?

1. **Dados de Giroscópio** - Estes dados devem retornar 3 valores (`x`, `y`, `z`). E devem ser armazenados juntamente com o `TIMESTAMP` do momento em que foi coletado;
2. **Dados de GPS** - Estes dados devem retornar 2 valores (`latitude` , `longitude`). E também devem ser armazenados juntamente com o `TIMESTAMP` do momento em que foram coletados;
3. **Uma foto** - Obter uma foto de uma das câmeras do dispositivo e enviá-la também junto com o `TIMESTAMP` em que foi coletada;

**🚨 É importante que se envie junto à essas informações um campo adicional, contendo uma identificação única do dispositivo, que pode ser seu endereço MAC.**

### Funcionamento

A aplicação Android deverá rodar em Background, e coletar e enviar as informações descritas a cada 10 segundos.

### Qual parte do desafio devo realizar?

Você deve realizar somente o desafio para a vaga que se candidatou.

Caso tenha sido a vaga de Android Embarcado, então resolva somente esta sessão.

Caso tenha sido a vaga de Backend, então resolva somente esta sessão.

---

# 🚀 Bora nessa!

Você deverá criar uma aplicação que irá receber os dados enviados pelo aplicativo.

Lembre-se essa aplicação precisa ser em GO!

## Nível 1 - API e Validação

Deve-se criar uma API que receba requisições de acordo com os endpoints:

- `POST /telemetry/gyroscope` - Dados do giroscópio;
- `POST /telemetry/gps` - Dados do GPS;
- `POST /telemetry/photo` - Dados da Foto;

Deve-se garantir que os dados recebidos estão preenchidos corretamente.

Caso algum dado esteja faltando, então retorne uma mensagem de erro e um Status 400.

## Nível 2 - Persistência e Testes

1. Salve cada uma das informações em um banco de dados a sua escolha.
2. Salve estes dados de forma identificável e consistente.
3. Crie testes unitários para cada arquivo da aplicação.
4. Implemente testes de integração para validar o fluxo completo.

## Nível 3 - Containerização e CI/CD

1. Crie um _container_ em _Docker_ que contenha a sua aplicação e o banco de dados.
2. Configure um pipeline de CI/CD (pode usar GitHub Actions, GitLab CI, ou similar).
3. Implemente testes automatizados no pipeline.
4. Configure o processo de build e deploy da aplicação.

## Nível 4 - Processamento de Imagens

1. Implemente a integração com AWS Rekognition para análise de fotos.
2. Compare cada nova foto com as fotos anteriores enviadas.
3. Retorne um atributo indicando se a foto foi reconhecida.
4. Implemente um sistema de cache para otimizar as consultas ao Rekognition.
5. Adicione logs detalhados do processo de análise.

## Nível 5 - Arquitetura com Filas

### Por que usar Filas?

Em um sistema de telemetria, o uso de filas é crucial por vários motivos:

1. **Desacoplamento**: Separa a coleta de dados do processamento, permitindo que cada parte evolua independentemente.
2. **Resiliência**: Se o processamento falhar, os dados não são perdidos, pois ficam na fila.
3. **Escalabilidade**: Permite processar mais dados adicionando mais consumidores.
4. **Pico de Carga**: Absorve picos de tráfego sem sobrecarregar o sistema.
5. **Processamento Assíncrono**: Permite que operações demoradas (como análise de imagens) não bloqueiem a coleta de dados.

### Implementação

- Implemente um sistema de mensageria (RabbitMQ, Apache Kafka, AWS SQS, NATS (usamos este)):
  - Crie filas separadas para cada tipo de telemetria
  - Implemente o padrão Producer-Consumer
  - Configure retry policies para mensagens com erro
  - Implemente dead letter queues para mensagens problemáticas

## Nível 6 - Monitoramento e Documentação

1. Implemente monitoramento e observabilidade:
   - Métricas de performance (Prometheus/Grafana)
   - Logs estruturados (ELK Stack ou similar)
   - Tracing distribuído
   - Alertas para condições anormais

2. Documentação Completa do Sistema:
   - Diagramas da arquitetura
   - Documentação técnica detalhada
   - Guia de operação e manutenção
   - Procedimentos de backup e recuperação
   - Guia de troubleshooting
   - Documentação de APIs usando OpenAPI/Swagger
   - Guia de segurança e boas práticas
   - Procedimentos de escalabilidade e resiliência
   - Documentação de configuração e variáveis de ambiente
   - Guia de contribuição e desenvolvimento

## Nível 7 - Segurança e Governança

Implemente camadas de segurança e governança:

1. Autenticação e autorização para todos os endpoints
2. Criptografia de dados sensíveis
3. Rate limiting para proteger a API
4. Validação de schema para todas as mensagens
5. Auditoria de todas as operações no sistema
