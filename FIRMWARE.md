# Desafio Firmware

## 💻 O Problema

Um dos nossos clientes ainda não consegue comprar o equipamento para colocar nos veículos de sua frota, mas ele quer muito utilizar a nossa solução.

Por isso, vamos fazer um MVP bastante simples para testar se, o celular do motorista poderia ser utilizado como o dispositivo de obtenção das informações.

> Parece fazer sentido certo? Ele possui vários mecanismos parecidos com o equipamento que oferecemos!

Sua missão ajudar na criação deste MVP para que possamos testar as frotas deste cliente.

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

## Nível 1

Deve-se coletar os dados de acordo com as especificações, e armazená-los em um banco de dados local;

## Nível 2

Deve-se criar testes unitários para garantir o funcionamento das estruturas criadas;

## Nível 3

Deve-se enviar os dados obtidos a cada 10 segundos para uma API com a seguinte rota

- `POST /telemetry/gyroscope` - Dados do giroscópio;
- `POST /telemetry/gps` - Dados do GPS;
- `POST /telemetry/photo` - Dados da Foto;

## Nível 4

Deve-se realizar um _crop_ da foto obtida para que se consiga extrair somente um rosto. Caso a foto não tenha um rosto, ela não deverá ser enviada.

## Nível 5

Faça com que cada uma das requisições ocorra de forma paralela, e não de forma síncrona;