# Desafio Firmware

## 💻 O Problema

Um dos nossos clientes ainda não consegue comprar o equipamento para colocar nos veículos de sua frota, mas ele quer muito utilizar a nossa solução.

Por isso, vamos fazer um MVP bastante simples para testar se o celular do motorista poderia ser utilizado como o dispositivo de obtenção das informações.

> Parece fazer sentido certo? Ele possui vários mecanismos parecidos com o equipamento que oferecemos!

Sua missão é ajudar na criação deste MVP para que possamos testar as frotas deste cliente.

Essa versão do produto será bastante simplificada. Queremos apenas criar as estruturas para obter algumas informações do seu dispositivo (Android) e armazená-la em um Banco de Dados.

Essas informações, depois de armazenadas devem estar disponíveis através de uma API para que este cliente integre com um Front-end já existente!

### Quais serão as informações que deverão ser coletadas?

1. **Dados de Giroscópio** - Estes dados devem retornar 3 valores (`x`, `y`, `z`). E devem ser armazenados juntamente com o `TIMESTAMP` do momento em que foi coletado;
2. **Dados de GPS** - Estes dados devem retornar 2 valores (`latitude` , `longitude`). E também devem ser armazenados juntamente com o `TIMESTAMP` do momento em que foram coletados;
3. **Uma foto** - Obter uma foto de uma das câmeras do dispositivo e enviá-la também junto com o `TIMESTAMP` em que foi coletada;

**🚨 É importante que se envie junto à essas informações um campo adicional, contendo uma identificação única do dispositivo, que pode ser seu endereço MAC.**

### Funcionamento

A aplicação Android deverá rodar em Background, e coletar e enviar as informações descritas a cada 10 segundos.

### Qual parte do desafio devo realizar?

Você deve realizar somente o desafio para a vaga que se candidatou.

Caso tenha sido a vaga de Android Embarcado, então resolva somente esta sessão.

Caso tenha sido a vaga de Backend, então resolva somente esta sessão.

---

# 🚀 Bora nessa!

Você deverá criar uma aplicação que deverá coletar os dados e enviá-los para o servidor Back-end;

Lembre-se que essa é uma aplicação Android nativa, e não deve possuir qualquer tipo de interface com o usuário.

## 📋 **Requisitos Técnicos Obrigatórios**

- **Java 17** - Linguagem principal
- **Android SDK** (API 21-34) - Foreground Services, Permissões
- **Banco de dados local** - Qualquer solução (Room, SQLite, ObjectBox, etc.)
- **Sistema de eventos/comunicação** - Qualquer solução (EventBus, RxJava, LiveData, etc.)
- **Multi-threading** - Threads, executors, sincronização
- **Testing** - Qualquer framework (JUnit, Robolectric, MockK, etc.)

---

## 🎯 **Níveis do Desafio**

### **Nível 1 - MVP Básico**
**Expectativa de Negócio:** Demonstrar que é possível coletar dados básicos do dispositivo e armazená-los localmente.

**Tarefas:**
1. Configurar projeto Android com Java 17
2. Implementar Foreground Service para coleta em background
3. Configurar permissões (LOCATION, CAMERA)
4. Criar entidades para giroscópio, GPS e foto
5. Implementar coleta básica dos sensores
6. Armazenar dados no banco local

---

### **Nível 2 - Arquitetura e Testes**
**Expectativa de Negócio:** Garantir que o sistema seja robusto, testável e mantenha qualidade em produção.

**Tarefas:**
1. **Implementar sistema de eventos/comunicação** entre componentes:
   - Criar eventos/mensagens para cada sensor
   - Usar o sistema escolhido para notificar quando novos dados são coletados
   - Implementar processamento assíncrono dos dados
   - **Objetivo:** Desacoplar componentes, facilitar testes, permitir processamento paralelo
2. Criar Repository Pattern para acesso aos dados
3. Implementar testes unitários
4. Adicionar testes de componentes Android
5. Implementar logs estruturados
6. Adicionar retry logic para falhas

---

### **Nível 3 - Comunicação e Performance**
**Expectativa de Negócio:** Permitir que os dados sejam enviados para o servidor de forma eficiente e confiável.

**Tarefas:**
1. Implementar upload HTTP para as APIs:
   - `POST /telemetry/gyroscope`
   - `POST /telemetry/gps`
   - `POST /telemetry/photo`
2. Implementar upload paralelo das requisições
3. Adicionar retry logic para falhas de rede
4. Otimizar uso de memória e CPU
5. Implementar compressão de dados

---

### **Nível 4 - Visão Computacional**
**Expectativa de Negócio:** Adicionar inteligência ao sistema para processar e validar imagens automaticamente.

**Tarefas:**
1. Integrar biblioteca de processamento de imagem (OpenCV, ML Kit, etc.)
2. Implementar detecção de rosto
3. Realizar crop automático da foto para extrair apenas o rosto
4. Só enviar fotos com rosto detectado

---

### **Nível 5 - Comunicação IoT e Serialização**
**Expectativa de Negócio:** Implementar comunicação em tempo real via MQTT e otimizar serialização de dados para IoT.

**Tarefas:**
1. **Implementar comunicação MQTT** com AWS IoT Core:
   - Configurar conexão MQTT segura com AWS
   - Criar tópicos para cada tipo de sensor
   - Implementar QoS adequado para cada tipo de dado
2. **Usar Protocol Buffers** para serialização:
   - Definir schemas .proto para cada tipo de telemetria
   - Serializar dados antes do envio MQTT
   - Implementar deserialização no lado servidor
3. **Manter comunicação HTTP** como fallback
4. **Implementar retry logic** específico para MQTT
5. **Otimizar payload** para reduzir uso de banda
6. **Implementar compressão** adicional se necessário

---

## 📊 **Critérios de Avaliação Geral**

### **Código e Arquitetura**
- Código limpo e bem estruturado
- Padrões arquiteturais adequados
- Separação de responsabilidades

### **Conhecimento Técnico**
- Java 17 e Android SDK
- Banco de dados e persistência
- Sistema de eventos e multi-threading
- Testing e debugging

### **Performance e Robustez**
- Otimização de memória e CPU
- Tratamento de erros robusto
- Concorrência e sincronização

---

## 🚀 **Bônus (Diferencial)**

- Implementar todos os 5 níveis
- Adicionar CI/CD
- Criar documentação técnica
- Implementar métricas de performance
- Adicionar sistema de configuração remota