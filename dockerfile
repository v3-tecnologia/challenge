# Usamos a imagem do Go como nossa imagem final.
# Ela é maior, mas contém todas as ferramentas que precisamos (como o 'go test').
FROM golang:1.24-alpine

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos de gerenciamento de dependências primeiro para cache
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o resto do código-fonte
COPY . .

# Constrói os DOIS binários (api e worker)
RUN go build -o /app/api ./cmd/api
RUN go build -o /app/worker ./cmd/worker

# Expõe a porta que nossa API usa
EXPOSE 8080

# Define o comando padrão para iniciar a API. O worker terá seu comando sobrescrito no docker-compose.
CMD ["/app/api"]