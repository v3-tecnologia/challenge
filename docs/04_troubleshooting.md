# Guia de Troubleshooting

**Projeto:** Solução de Telemetria de Frota (MVP)  
**Versão:** 1.0

Este documento serve como um guia de primeiros socorros para diagnosticar e resolver os problemas mais comuns em ambiente de desenvolvimento.

---

## Problema 1: Aplicação (app ou worker) não inicia e o log mostra `connection refused`

**Sintoma:**  
Erro como: `dial tcp [::1]:5432: connect: connection refused`.

**Causa Provável:**  
A aplicação tentou se conectar ao banco, mas ele não estava disponível.

**Soluções:**

- Verifique se o container do banco está rodando:

```bash
docker-compose ps
```

O serviço `challenge_db_postgres` deve estar com status `Up`. Caso contrário:

```bash
docker-compose up -d db
```

- Verifique o mapeamento de portas (porta 5432 deve estar exposta).
- Desative temporariamente o firewall ou antivírus que possa estar bloqueando a porta.

---

## Problema 2: Aplicação não inicia e o log mostra `no such host`

**Sintoma:**  
Erro como: `dial tcp: lookup db: no such host`.

**Causa Provável:**  
A configuração do `DB_HOST` não está compatível com o ambiente.

**Soluções:**

- Se estiver rodando com `go run main.go`, use:

```env
DB_HOST=localhost
```

- Se estiver rodando com `docker-compose up`, use:

```env
DB_HOST=db
```

---

## Problema 3: Worker mostra erro da AWS `InvalidSignatureException`

**Sintoma:**  
Erro como: `InvalidSignatureException: The request signature we calculated does not match...`.

**Causa Provável:**  
Credenciais AWS incorretas.

**Soluções:**

- Verifique as variáveis `AWS_ACCESS_KEY_ID` e `AWS_SECRET_ACCESS_KEY` no `.env`.
- Remova espaços em branco no início/fim das chaves.
- Gere novas credenciais no console AWS (IAM).
- Reinicie os serviços:

```bash
docker-compose down
docker-compose up --build
```

---

## Problema 4: Worker mostra erro da AWS `ResourceNotFoundException`

**Sintoma:**  
Erro como: `The collection id: fleet_drivers_test does not exist`.

**Causa Provável:**  
A coleção do Rekognition ainda não foi criada no ambiente.

**Solução:**  
Certifique-se de que a aplicação executa `rekognitionClient.CreateCollection(...)` durante a inicialização, com tratamento de erro para quando a coleção já existir.

---

## Problema 5: Build do Docker falha com erro `snapshot not found`

**Sintoma:**  
Erro como: `parent snapshot ... does not exist` ou `failed to resolve source metadata`.

**Causa Provável:**  
Cache de build do Docker corrompido.

**Soluções:**

```bash
docker-compose down

docker builder prune -a -f

docker-compose up --build
```

## Problema 6: API retorna erro 401 Unauthorized

**Sintoma:**  
Ao enviar uma requisição para um endpoint de telemetria, a resposta é `401 Unauthorized`.

**Causa Provável:**  
A Chave de API (API Key) está faltando ou está incorreta.

**Soluções:**

- Verifique se sua requisição (ex: Postman) inclui o cabeçalho HTTP `X-API-Key`.
- Confirme que o valor enviado nesse cabeçalho é exatamente o mesmo definido na variável `API_KEY` do seu arquivo `.env`.

---

## Problema 7: API retorna erro 429 Too Many Requests

**Sintoma:**  
Após enviar várias requisições rapidamente, a API começa a responder com `429 Too Many Requests`.

**Causa Provável:**  
O Rate Limiter foi ativado e o endereço IP excedeu o limite de requisições por segundo.

**Soluções:**

- Esse é o comportamento esperado. A API está se protegendo contra abuso.
- Espere um ou dois segundos para o "balde de fichas" do seu IP ser reabastecido e tente novamente.
- Para desativar temporariamente durante testes de carga, comente a linha do `RateLimiterMiddleware` no `cmd/api/main.go`.

---