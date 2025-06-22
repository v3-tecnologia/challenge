<h1 align="center" style="font-weight: bold;">Decisões</h1>

## Pensamentos

Esse arquivo é referente aos pensamentos, ideias e decisões. Como cheguei a alguma decisão, o desenvolvimento de funcionalidades e etc. Digamos que é um segundo cerebro, quem sabe essa anotação vira um artigo no medium.


## Decisões Iniciais
Para inicio decidir a arquitertura é importante para termos organização e após pensar e ler alguns artigos decidi que a melhor escolha seria a junção de clean arquiterture e hexagonal 

https://dev.to/espigah/go-estrutura-de-projetos-1j0k<br>
https://johnfercher.medium.com/go-arquitetura-hexagonal-dbcd2e986b55<br>
https://dev.to/booscaaa/implementando-clean-architecture-com-golang-4n0a

## Nível 1
Um ponto relevante foi a validação, tive que pesquisar e o giroscópio pode retornar 0 que seria o norte além da latitude e longitude então a utilização do * no float64 para comparar se é nulo

## Nível 2
Comecei Fazendo os primeiros testes e esse artigo é ótimo para se basear, interessante adicionar o -cover para ser possível analisar a cobertura de testes
https://medium.com/@habbema/golang-testes-86da3e5e0602