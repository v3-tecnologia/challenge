# Desafio Android

## Estratégia do nível 1

* Escolher banco de dados: Room
* Criar um banco local
* Simular a coleta de dados
* Gravar dandos no banco
* Garantir que a aplicação capte os dados de 10 em 10 segundos

### Banco Room

É um banco prático e leve, de alta performance. Para colocá-lo na aplicação, basta acrescentá-lo à dependência do gradle

### Criação do banco e entidades

Devem ser criadas as entidades e depois os Dao para fazer o CRUD. Não acrescentei Update por não ver muito sentido em alterar um dado que é captado. Utilizei o Volatile e Synchronized para garantir uma instância só do banco e evitar problemas com corrida de threads.

### Gravação de dados

Acontece na classe DataCollectService. Os dados foram simulados. Alguns métodos, mesmo não sendo abstrados, foram implementados por meio de polimorfismo, pois foi necessário adaptar para os requisitos do desafio (10 em 10 segundos)

### 10 em 10 segundos

Método na MainActivity que realiza a ação principal.

## Estratégia do nível 2

O Android vem com JUnit 4 como padrão, mas optei pelo JUnit5 por questões de conforto. Já tenho o costume de usar JUnit 5 com Mockito por vir atuando com Spring Boot recentemente com versões do Java mais atuais (da 17 a superiores). Tive um problema com o SystemClock do Android. Acabei deixando esse passo parado, pois o teste será alterado com as etapas seguintes, então não faz sentido antecipar o teste. Teoricamente acho mais correto o teste ser a etapa 5. 

## Estratégia do nível 4

Optei pela biblioteca JavaCV. Eu conheço duas técnicas de reconhecimento facial, a Eigenfaces e a JavaCV(que não é muito diferente do OpenCV no Python). Me senti mais capaz de utilizar a JavaCV. No caso, já tinha alguns códigos prontos da minha pós-graduação, então apenas precisei adaptar para o projeto. Como é uma alteração que envolve uma classe apenas (a de fotos), também achei mais prático vir para o nível 4, deixando o 3 pra ser realizado posteriormente (pois alterará os testes, quando vi o nível 2, ainda não havia pensado no 3). Utilizei um modelo pronto xml do Java. Creio que a proposta não seja entrar nos méritos mais aprofundados de machine learning (redes neurais convolucionais, etc), portanto optei por não usar treinamentos, etc. Criei uma classe FaceDetector bem simples, utilizando um modelo já treinado e que é bastante comum em aplicações. 