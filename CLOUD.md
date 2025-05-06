# Desafio Cloud

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

Você deverá criar uma aplicação que irá receber os dados enviados pelo aplicativo.

Lembre-se essa aplicação precisa ser em GO!

## Nível 1

Deve-se criar uma API que receba requisições de acordo com os endpoints:

- `POST /telemetry/gyroscope` - Dados do giroscópio;
- `POST /telemetry/gps` - Dados do GPS;
- `POST /telemetry/photo` - Dados da Foto;

Deve-se garantir que os dados recebidos estão preenchidos corretamente.

Caso algum dado esteja faltando, então retorne uma mensagem de erro e um Status 400.

## Nível 2

Salve cada uma das informações em um banco de dados a sua escolha.

Salve estes dados de forma identificável e consistente;

## Nível 3

Crie testes unitários para cada arquivo da aplicação. Para cada nova implementação a seguir, também deve-se criar os testes.

## Nível 4

Crie um _container_ em _Docker_ que contenha a sua aplicação e o banco de dados utilizado nos testes.

## Nível 5

A cada foto recebida, deve-se utilizar o AWS Rekognition para comparar se a foto enviada é reconhecida com base nas fotos anteriores enviadas.

Se a foto enviada for reconhecida, retorne como resposta do `POST` um atributo que indique isso.

Utilize as fotos iniciais para realizar o treinamento da IA.
