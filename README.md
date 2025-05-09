<h1 align="center" style="font-weight: bold;"> Desafio BackEnd V3 </h1>

### 🛠 Tecnologias Utilizadas
- **Go** 1.24.2  
- **Docker**  
- **MySQL**

---

## 🚀 Começando <a id="getting-started"></a>

### ✅ Pré-requisitos
- [Docker](https://www.docker.com/) instalado na sua máquina.

---

### ⚙️ Configuração do Ambiente

Renomeie o arquivo `.env.example` para `.env` e preencha as variáveis:

| Variável       | Descrição                                                       |
|----------------|-----------------------------------------------------------------|
| `DB_HOST`      | Nome do serviço MySQL no Docker ou endereço da máquina         |
| `DB_PORT`      | Porta do MySQL (padrão: `3306`)                                 |
| `DB_USER`      | Usuário do banco de dados                                       |
| `DB_PASSWORD`  | Senha do banco de dados                                         |
| `DB_NAME`      | Nome do banco de dados                                          |
| `DB_LOG`       | Ativa log do banco (`1` para ativar, `0` para desativar`)       |

---

### ▶️ Executando a Aplicação

Para iniciar a aplicação e o banco de dados, utilize:

```bash
sudo docker-compose up --build
```

### 🧪 Executando os Testes

Para rodar os testes, utilize:

```bash
sudo docker-compose run --rm test
```

