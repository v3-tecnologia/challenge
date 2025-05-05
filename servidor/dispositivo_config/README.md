# 📡 V3 Dispositivos - Servidor

Este projeto implementa um servidor HTTP assíncrono usando [FastAPI](https://fastapi.tiangolo.com/) que expõe o endpoint `/api/dispositivos/configurar`.  
O objetivo é receber comandos de configuração de dispositivos embarcados, como o ajuste de volume de alerta para veículos.

## 🚀 Funcionalidades

- Endpoint `POST /api/dispositivos/configurar`
- Validação de payloads JSON contendo:
  - `placa` (string): placa do veículo
  - `volume_alerta` (inteiro de 0 a 100): volume desejado

- Resposta com status `200 Ok` em caso de sucesso
- Projeto estruturado com [Poetry](https://python-poetry.org/) para gerenciamento de dependências
- Suporte a testes unitários com `pytest`
- Código formatado e validado com `black` e `flake8`
- Tarefas automatizadas via `Makefile`

---

## 🧰 Requisitos

Antes de executar o projeto, certifique-se de ter os seguintes softwares instalados no seu ambiente:

- [Python 3.10+](https://www.python.org/downloads/)
- [Poetry](https://python-poetry.org/docs/#installation)
- [Make (GNU Make)](https://www.gnu.org/software/make/) (Linux/macOS já possuem, no Windows use Git Bash ou WSL)

---

## ⚙️ Instalação

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/v3-dispositivos-servidor.git
cd v3-dispositivos-servidor
``` 

2. Instale as dependências do projeto:

```bash
make install
``` 

## 🛠️ Comandos Disponíveis

Este projeto utiliza make para facilitar tarefas administrativas. Veja os comandos disponíveis:

Comando          | Descrição                                     |
---------------- | --------------------------------------------- |
`make install`   | Instala as dependências via Poetry            |
`make build`	 | Prepara o ambiente (lint + format)            |
`make run`       | Inicia o servidor FastAPI local na porta 8000 |
`make test`      | Executa os testes unitários com pytest        |
`make format`    | Formata o código com Black                    |
`make lint`      | Verifica padrões de código com flake8         |
`make clean`     | Remove arquivos temporários                   |

## 🔥 Executando a API
Após instalar tudo, inicie o servidor com:

```bash
make run
``` 

A API estará disponível em:
📍 http://localhost:8000/api/dispositivos/configurar

## 📬 Exemplo de Requisição

```bash
POST /api/dispositivos/configurar
Content-Type: application/json

{
  "placa": "ABC1234",
  "volume_alerta": 50
}

``` 

Ou utilizando nossos scripts:

```bash
/ajustar_volume.sh -f veiculos-ok.csv -u http://localhost:8000/api/dispositivos/configurar
```

## 🧪 Testes

```bash
make test
```
