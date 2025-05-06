### Stack utilizada

- Linguagem: Go (Golang)
- Armazenamento: PostgreSQL com sqlc
- Arquitetura: Este projeto segue uma arquitetura **hexagonal (ports & adapters)** com forte separação de responsabilidades, visando testes fáceis, baixo acoplamento e extensibilidade futura.
- Testes: testify, unitários para entidades e usecases
- Storage de imagem: local (./uploads) via interface desacoplada
- Reconhecimento facial: mock (InMemoryFaceRecognition)
- Deploy local: Docker

---

### Requisitos e Níveis

| Nível | Descrição                                                             | Status |
| ----- | --------------------------------------------------------------------- | ------ |
| 1     | Criar API com endpoints `POST /telemetry/gps`, `/gyroscope`, `/photo` | ✔️     |
|       | Validar campos obrigatórios (device_id, timestamp, etc.)              | ✔️     |
| 2     | Persistir os dados em banco de dados                                  | ✔️     |
|       | Utilizei PostgreSQL + sqlc, com migrations                            | ✔️     |
| 3     | Testes unitários para cada arquivo da aplicação                       | ✔️     |
|       | Testes em entidades, usecases e validações com `testify`              | ✔️     |
| 4     | Docker com banco de dados e aplicação Go                              | ✔️     |
|       | Makefile para rodar testes, gerar sqlc, criar migrations              | ✔️     |
| 5     | Integração com AWS Rekognition                                        | ❌     |
|       | (Atualmente usando mock `InMemoryFaceRecognition`)                    | ✖️     |

---

### Detalhes adicionais implementados

- Validação com erros customizados (appError)
- Upload de imagem com multipart/form-data
- Armazenamento de imagem via interface PhotoStorageService
- Modularização por pacotes (entity, repository, usecase, controller, route)
- Documentação de API com exemplos cURL
- Rotas organizadas com chi
- Makefile organizado com comandos úteis
- .gitignore configurado para ignorar uploads

---

### Melhorias futuras (caso evolua o MVP)

- Substituir armazenamento local por S3PhotoStorage
- Substituir mock de Rekognition por integração real via AWS SDK
- Processar reconhecimento facial de forma assíncrona (via fila, ex: SQS)
