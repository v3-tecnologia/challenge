# Guia de Segurança e Boas Práticas

**Projeto:** Solução de Telemetria de Frota (MVP)
**Versão:** 1.0

## 1. Visão Geral
Este documento descreve a postura de segurança da aplicação, os controles implementados e as boas práticas que devem ser seguidas por desenvolvedores e operadores para garantir a integridade, confidencialidade e disponibilidade do sistema.

## 2. Modelo de Ameaças Consideradas
A arquitetura atual implementa defesas contra as seguintes ameaças principais:
- **Acesso Não Autorizado à API:** Tentativas de envio de dados por clientes não autorizados.
- **Abuso de API / Negação de Serviço (DoS):** Clientes mal-intencionados ou com bugs enviando um volume excessivo de requisições para degradar o serviço.
- **Exposição de Dados Sensíveis em Repouso:** Risco de vazamento de dados caso o banco de dados seja comprometido.
- **Exposição de Segredos de Configuração:** Risco de chaves de API, senhas e outras credenciais serem expostas no código-fonte.

## 3. Controles de Segurança Implementados

### 3.1. Autenticação de API
- **Mecanismo:** Autenticação baseada em Chave de API (API Key).
- **Implementação:** Todas as requisições para os endpoints de telemetria (`/telemetry/*`) devem incluir o cabeçalho HTTP `X-API-Key` contendo um token secreto pré-definido. Um middleware na API valida esta chave. Requisições sem a chave ou com uma chave inválida são rejeitadas com `HTTP 401 Unauthorized`.

### 3.2. Rate Limiting (Controle de Taxa de Requisições)
- **Mecanismo:** Limitação de taxa por endereço de IP.
- **Implementação:** Um middleware na API controla o número de requisições que cada IP pode fazer por segundo. Por padrão, o limite é de 5 requisições/segundo com um pico permitido de 10. Se um cliente exceder este limite, ele receberá uma resposta `HTTP 429 Too Many Requests`. Isso protege a API contra sobrecarga.

### 3.3. Criptografia de Dados em Repouso
- **Mecanismo:** Criptografia simétrica AES-256-GCM.
- **Implementação:** O dado mais sensível, a `photo`, é criptografado pelo `worker` **antes** de ser salvo no banco de dados PostgreSQL. Isso garante que, mesmo com acesso direto ao banco, o dado da imagem não pode ser lido sem a chave de criptografia.

### 3.4. Gestão de Segredos
- **Mecanismo:** Variáveis de ambiente carregadas a partir de um arquivo `.env`.
- **Implementação:** Todas as informações sensíveis são definidas no arquivo `.env`, que é explicitamente ignorado pelo Git (`.gitignore`). No ambiente de CI/CD, esses valores são injetados de forma segura através dos **GitHub Secrets**.

## 4. Boas Práticas de Segurança Recomendadas

- **Princípio do Menor Privilégio:** O usuário IAM configurado na AWS possui apenas as permissões estritamente necessárias para a aplicação funcionar (`AmazonRekognitionFullAccess`).
- **HTTPS em Produção:** Em um ambiente de produção, a API deve ser exposta exclusivamente via HTTPS para garantir a criptografia em trânsito. Isso geralmente é configurado em um Load Balancer ou Reverse Proxy na frente da aplicação.
