## Cloud Telemetry API

Este repositório contém a implementação de um MVP para coleta de dados de telemetria (GPS, giroscópio e fotos) de motoristas utilizando seus próprios dispositivos Android.

A API foi construída em **Go** seguindo princípios de arquitetura limpa, com foco em **desacoplamento**, **testes** e **validações** desde o início.

---

### Como rodar localmente

```bash
docker-compose up --build
```

---

### Funcionalidades implementadas

- ✅ API REST com endpoints para GPS, Giroscópio e Foto
- ✅ Validação de campos obrigatórios (device_id, timestamp, etc.)
- ✅ Armazenamento persistente com PostgreSQL via `sqlc`
- ✅ Upload de imagem via `multipart/form-data`
- ✅ Mock de reconhecimento facial (modo local)
- ✅ Testes unitários para entidades e regras de negócio
- ✅ Dockerfile e Makefile para facilitar setup local

---

### Documentação

A pasta [`/docs`](./docs) contém arquivos que detalham cada aspecto do projeto:

| Documento                                                                 | Descrição                               |
| ------------------------------------------------------------------------- | --------------------------------------- |
| [`endpoints.md`](./docs/endpoints.md)                                     | Exemplos de uso da API com `cURL`       |
| [`requisitos-para-desenvolver.md`](./docs/requisitos-para-desenvolver.md) | Ferramentas e configurações necessárias |
| [`status-de-implementacao.md`](./docs/status-de-implementacao.md)         | Checklist do desafio e nível de entrega |

---

### ⚠️ Observação

A integração real com AWS Rekognition não foi implementada. Atualmente, o reconhecimento facial é simulado via InMemoryFaceRecognition.
