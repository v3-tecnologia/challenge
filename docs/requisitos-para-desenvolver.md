### Requisitos Gerais

- Sistema: macOS ou Linux (preferencialmente Unix-like)
- Terminal com acesso a `make`
- Conexão com internet para instalar ferramentas

---

### Ferramentas Necessárias

- `Go`: Versão preferível 1.20+

- `Sqlc`:
  - Utilizado para gerar código Go a partir de queries SQL
  - Gera interfaces seguras, structs e execuções baseadas em parâmetros

```bash
brew install sqlc
# Ou
# curl -sSfL https://github.com/sqlc-dev/sqlc/releases/latest/download/sqlc_$(uname -s)_amd64.tar.gz | tar -xz -C /usr/local/bin
```

> Verifique: `sqlc version`

---

### 3. **migrate (golang-migrate)**

- Gerencia migrações do banco de dados via CLI
- Criação e execução de migrations com versionamento

```bash
brew install golang-migrate
# Ou via script Linux
# curl -L https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
# sudo apt install migrate
```

> Verifique: `migrate -version`

---

### 4. **Docker**

- Para rodar banco de dados localmente

> Verifique: `docker --version`

---

### Comandos via `make`

| Comando                 | Descrição                                 |
| ----------------------- | ----------------------------------------- |
| `make test`             | Roda todos os testes unitários            |
| `make create_migration` | Cria nova migration SQL via CLI           |
| `make generate_sqlc`    | Gera o código Go com base nas queries SQL |

---

### Banco de Dados

- Recomendado: PostgreSQL 15+
- Rodar localmente com Docker (via `docker-compose.yml`)
- Migrations são aplicadas via `migrate`
