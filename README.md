# Resposta de Desafio Técnico

<p align="center">
    <img src="./.github/logo.png" width="200px" alt="Logo do Projeto">
</p>

## Requisitos

- **Python 3.12** ou superior
- **pip** (gerenciador de pacotes Python)
- **virtualenv** ou **venv** para ambiente virtual

---

## Configuração do Ambiente

### 1. Clone o repositório
```bash
git clone <url-do-repositorio>
cd <nome-do-repositorio>
```

### 2. Criação do ambiente virtual

#### Linux/macOS
```bash
python3 -m venv venv
source venv/bin/activate
```

#### Windows
```bash
python3 -m venv venv
venv\Scripts\activate
```

### 3. Instalação das dependências
```bash
pip install -r requirements.txt
```

### 4. Configuração do banco de dados
```bash
python manage.py migrate
```

### 5. Inicialização do servidor
```bash
python manage.py runserver
```

---

## Documentação da API

Após iniciar o servidor, acesse a documentação interativa da API disponivel em `http://ip_da_aplicação:8000/swagger/`
substitua ip_da_aplicação pelo ip que está executando o servidor, ou use localhost para acessar-lo a partir da sua propria maquina

---

## Notas

- Certifique-se de que o ambiente virtual está ativado antes de executar os comandos
- O arquivo `requirements.txt` contém todas as dependências necessárias
- Para parar o servidor, use `Ctrl+C` no terminal
- 
---

## Agradecimentos
Gostaria de agradecer à V3 pela oportunidade de demonstrar minhas habilidades técnicas através deste desafio.
Espero muito conseguir avançar no processo seletivo e contribuir com meu conhecimento para esta incrível empresa.
