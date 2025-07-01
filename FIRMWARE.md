# Desafio Firmware

Desafio: MVP para Coleta de Dados de Dispositivo Móvel

## 📌 Contexto do Negócio

Precisamos validar se smartphones Android podem substituir dispositivos IoT na coleta de dados de frotas. Seu objetivo é criar uma solução que:

* Rode continuamente em segundo plano, sem interface gráfica;

* Capture giroscópio (x,y,z), GPS (lat/long) e fotos a cada 10 segundos, ou baseado em eventos (Ex: Quando a posição mudar);

* Associe um identificador único do dispositivo (ex: MAC);

* Armazene e disponibilize os dados para integração

* Capture uma foto do Motorista quando o sistema for acionado

* A Nuvem deve ser capaz de se comunicar com todos ou com um dispositivo específico para o envio de configuração e/ou comandos de qualquer tipo (***Diferencial***) (Ex: Mudar o tempo de envio das informações, mudar regra de envio de posicionamento, desativar um recurso, etc...)

---

## ✅ Requisitos Essenciais

1. **Coleta Contínua**

    * Funcionamento em background com intervalo fixo de 10s

    * Resiliência contra interrupções do sistema

2. **Dados Obrigatórios**

    * Giroscópio + timestamp

    * Coordenadas GPS + timestamp

    * Foto + timestamp

    * Identificador único do dispositivo

3. **Armazenamento**

    * Persistência local confiável

    * Recuperação segura após falhas

4. **Disponibilização**

    * Os dados armazenados devem ser enviados para a Nuvem de maneira confiável

    * Como se trata de um dispositivo embarcado a comunicação deve ser performática

    * Escolha o melhor protocolo ou tipo integração para atender aos requisitos acima

    * Todos os dados obrigatórios após coletados devem ser enviados em um mesmo payload. 
      Deve-se garantir que os dados enviados se refiram ao mesmo momento (exige coordenação dos eventos)

--- 

## Requisitos de Tecnologia 

* Java/Kotlin
* Android SDK (API 21-34) - Foreground Services, Permissões
* Banco de dados local - Qualquer solução (Room, SQLite, ObjectBox, etc.)
* Pense fora da caixa (Modelo de Atores, Coroutines, EDA, gRPc, etc...)

---

## 🎯 **Níveis do Desafio**

### Nível 1: Validação do Conceito
- Prove que os dados podem ser coletados e persistidos localmente
- Garanta funcionamento contínuo em segundo plano

### Nível 2: Confiabilidade Industrial
- Implemente tratamento de erros robusto
- Adicione testes automatizados críticos
- Documente estratégias de recuperação de falhas

### Nível 3: Preparação para Escala
- Crie um mecanismo de envio dos dados para a Nuvem
- Otimize consumo de recursos (bateria, rede)
- Implemente segurança básica na comunicação

### Nível 4: Inteligência Embarcada (***Diferencial***)
- Adicione validação automática de fotos (Ex: Validar se há um rosto em frente a camera)
- Implemente filtros para dados inconsistentes (Ex: Lat/Long zerados)

### Nível 5: Otimização Avançada (***Diferencial***)
- Explore protocolos alternativos para envio eficiente
- Implemente serialização otimizada para IoT

---

## 📊 **Critérios de Avaliação Geral**

### **Código e Arquitetura**
- Código limpo e bem estruturado
- Padrões arquiteturais adequados
- Separação de responsabilidades

### **Performance e Robustez**
- Otimização de memória e CPU
- Tratamento de erros robusto
- Concorrência e sincronização

---

***Boa sorte <3***