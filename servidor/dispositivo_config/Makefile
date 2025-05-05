# Nome do serviço
APP_NAME = dispositivo_config

# Comando padrão do Poetry
POETRY = poetry

# Arquivo principal da aplicação FastAPI
ENTRYPOINT = dispositivo_config.main:app

# Porta padrão
PORT = 8000

# ----------------------------------------
# Comandos
# ----------------------------------------

.PHONY: help install build run test lint format clean

help:
	@echo "Comandos disponíveis:"
	@echo "  make install       Instala dependências com poetry"
	@echo "  make build         Prepara ambiente (instala dependências + formata + lint)"
	@echo "  make run           Executa o servidor FastAPI com uvicorn"
	@echo "  make test          Executa os testes unitários com pytest"
	@echo "  make lint          Verifica formatação e boas práticas com flake8"
	@echo "  make format        Formata o código com black"
	@echo "  make clean         Remove arquivos temporários"

install:
	$(POETRY) install

build: install format lint

run:
	$(POETRY) run uvicorn $(ENTRYPOINT) --reload --port $(PORT)

test:
	$(POETRY) run pytest -v

lint:
	$(POETRY) run flake8 $(APP_NAME)

format:
	$(POETRY) run black $(APP_NAME) tests

clean:
	find . -type d -name "__pycache__" -exec rm -r {} +
	find . -type f -name "*.pyc" -delete
