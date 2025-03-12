# Desafio Técnico V3

## ❤️ Bem vindos

Olá, tudo certo?

Seja bem vindo ao teste de seleção para novos desenvolvedores na V3!

Estamos honrados que você tenha chegado até aqui!

Prepare aquele ☕️ , e venha conosco codar e se divertir!

## Poxa, outro teste?

Nós sabemos que os processos de seleção podem ser ingratos! Você investe um tempão e no final pode não ser aprovado!

Aqui, nós presamos pela **transparência**!

Este teste tem um **propósito** bastante simples:

> Nós queremos avaliar como você consegue transformar problemas em soluções através de código!

**🚨 IMPORTANTE!** Se você entende que já possui algum projeto pessoal, ou contribuição em um projeto _open-source_ que contemple conhecimentos equivalentes aos que existem neste v3pb-desafio, então, basta submeter o repositório explicando essa correlação!

## 🚀 Bora nessa!

Este é um teste para analisarmos como você desempenha ao entender, traduzir, resolver e entregar um código que resolve um problema.

### Dicas

- Documente seu projeto;
- Faça perguntas sobre os pontos que não ficaram claros para você;
- Mostre a sua linha de raciocínio;
- Trabalhe bem o seu README.md;
  - Explique até onde implementou;
  - Como o projeto pode ser executado;
  - Como pode-se testar o projeto;

### Como você deverá desenvolver?

1. Faça um _fork_ deste projeto em seu GitHub pessoal;
2. Realize as implementações de acordo com cada um dos níveis;
3. Faça pequenos _commits_;
4. Depois de sentir que fez o seu máximo, faça um PR para o repositório original.

🚨 **IMPORTANTE!** Não significa que você precisa implementar **todos os níveis** para ser aprovado no processo! Faça até onde se sentir confortável.

### Qual o tempo para entregar?

Quanto antes você enviar, mais cuidado podemos ter na revisão do seu teste. Mas sabemos que o dia a dia é corrido, faça de forma que fique confortável para você!

**Mas não desista! Envie até onde conseguir.**

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

### Qual parte do v3pb-desafio devo realizar?

Você deve realizar somente o v3pb-desafio para a vaga que se candidatou.

Caso tenha sido a vaga de Android Embarcado, então resolva somente esta sessão.

Caso tenha sido a vaga de Backend, então resolva somente esta sessão.

---

# Desafio Android Embarcado

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

# Desafio Backend

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
