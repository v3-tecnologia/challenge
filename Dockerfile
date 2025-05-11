# FROM postgres:latest

# # Define variáveis de ambiente
# ENV POSTGRES_USER=meuusuario
# ENV POSTGRES_PASSWORD=minhasenha
# ENV POSTGRES_DB=meubanco

# # Copia arquivos .sql para execução na inicialização
# COPY init.sql /docker-entrypoint-initdb.d/
# Use a imagem oficial do Golang
FROM golang:1.24.3

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie os arquivos Go para o diretório de trabalho
COPY . .

# Baixe as dependências do Go
RUN go mod tidy

# Compile a aplicação (opcional, você pode usar go run diretamente)
# RUN go build -o main .

# Exponha a porta 8080 para comunicação externa
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["go", "run", "cmd/main.go"]
