## O sistema é composto pelos seguintes serviços, orquestrados via docker-compose:

### 3.1. API (Producer)

- **Repositório:** `challenge-v3`  
- **Container:** `challenge_app_api`  
- **Tecnologia:** Go (`golang:1.24-alpine`)  
- **Responsabilidade:**  
  É o ponto de entrada (gateway) para todos os dados de telemetria. Sua responsabilidade principal é receber as requisições HTTP e publicá-las na fila NATS o mais rápido possível. Além disso, atua como a primeira linha de defesa do sistema, aplicando uma cadeia de middlewares a todas as requisições de telemetria para garantir Autenticação (via X-API-Key), Rate Limiting (por IP) e a coleta de Métricas (para o Prometheus).  

- **Endpoints:**  
  - `POST /telemetry/gyroscope`  
  - `POST /telemetry/gps`  
  - `POST /telemetry/photo`  

- **Comunicação:**  
  Recebe requisições HTTP da internet, publica mensagens para o serviço NATS.

---

### 3.2. Worker (Consumer)

- **Repositório:** `challenge-v3`  
- **Container:** `challenge_app_worker`  
- **Tecnologia:** Go (`golang:1.24-alpine`)  
- **Responsabilidade:**  
  Realiza o processamento pesado e assíncrono das mensagens. Ele se inscreve nos tópicos do NATS, consome as mensagens uma a uma e executa a lógica de negócio principal.

  Para mensagens de foto, ele interage com o AWS Rekognition e gerencia um cache em memória. Antes de persistir os dados, ele criptografa o dado da foto (usando AES-GCM) para garantir a segurança em repouso.

  Para todos os tipos de telemetria, o worker registra um evento de auditoria no banco de dados após cada processamento bem-sucedido.

- **Comunicação:**  
  Consome mensagens do serviço NATS, envia requisições para a API externa AWS Rekognition e escreve no serviço DB (PostgreSQL).

---

### 3.3. NATS JetStream (Fila de Mensagens)

- **Container:** `challenge_nats`  
- **Tecnologia:** `nats:2.10-alpine` com JetStream (`-js`) ativado  
- **Responsabilidade:**  
  Atua como um buffer resiliente e persistente entre a API e o Worker. Garante que nenhuma mensagem seja perdida, mesmo que o Worker esteja offline ou falhe.  

  Suporta políticas de reentrega (`Nack`) e descarte de mensagens problemáticas para uma *Dead-Letter Queue* (`Term`).  

- **Configuração:**  
  Utiliza um Stream chamado `TELEMETRY` que captura todos os subjects no padrão `telemetry.*`.

---

### 3.4. PostgreSQL (Banco de Dados)

- **Container:** `challenge_db_postgres`  
- **Tecnologia:** `postgres:15`  
- **Responsabilidade:**  
  Armazenamento persistente e relacional de todos os dados de telemetria que foram processados com sucesso pelo Worker.  

- **Schema:**  
  Contém 4 tabelas principais: `gyroscope`, `gps`, `photo` e `audit_log` para registrar as operações do sistema. Cada tabela possui colunas bem definidas para garantir a consistência dos dados.

---

### 3.5. AWS Rekognition (Serviço Externo)

- **Tecnologia:** Serviço gerenciado da AWS (SaaS)  
- **Responsabilidade:**  
  Fornece a funcionalidade de Inteligência Artificial para análise de imagens e reconhecimento facial.  

- **Interação:**  
  O Worker usa a API `SearchFacesByImage` para comparar um rosto com uma coleção pré-existente e `IndexFaces` para adicionar novos rostos a essa coleção.  

  A coleção usada é definida pela variável de ambiente `REKOGNITION_COLLECTION_ID`.

---

## 4. Estrutura do Projeto

O código-fonte está organizado nos seguintes pacotes principais:

- `cmd/`: Contém os pontos de entrada para os binários compiláveis (`api` e `worker`)  
- `handlers/`: Lógica da camada de API, responsável por lidar com as requisições HTTP  
- `services/`: Contém a lógica de negócio principal (ex: `PhotoAnalyzerService`)  
- `storage/`: Camada de acesso a dados, responsável pela comunicação com o banco de dados  
- `models/`: Definição das estruturas de dados (`structs`) e suas validações  
- `messaging/`: Funções auxiliares para conexão e configuração do NATS
- `metrics/`: Definição e exposição das métricas para o Prometheus
- `crypto/`: Definição e exposição das métricas para o Prometheus.  
- `ierr/`: Definição de tipos de erro personalizados
