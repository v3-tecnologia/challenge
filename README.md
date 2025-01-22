# Desafio V3

## Elucidação sobre o Nível 2

Sobre a necessidade de testes unitários:
Testes unitários não são rigorosamentes necessários onde não existém
processamentos de dados em partes particulares do projeto.
Em procedimentos onde existém apenas implementações para a obtenção e
transferência de informações, garantir a consistência e integridade,
são fatores fundamentais para o bom funcionamento desta parte particular do projeto.

## Documentação interna

Descrições através de comentários estão expostos
ao longo do projeto para uma elucidação sobre a arquitetura.

## Abstrações técnicas

Inibindo em algumas áreas implementações que sériam puramente técnicas,
onde não existem proporções algorítmicas a serem resolutas.
Está realizada a abstração arquitetural, inibindo algumas implementaçẽos
puramente técnicas para esta resolução. Poderão ser observadas aos
comentários.


## Implementações

Muitas implementações são de natureza arquiteturais, inibindo muitos outros detalhes de natureza puramente técnica na execução de determinados procedimentos.

- implementação de Mediator de serviços de controlador de procedimentos concorrentes, e finalizador dos procedimentos concorrentes;
- implementação de Decorator para procedimentos de requisições HTTP, e envio de requisições concorrentes;
- implementação de Service para procedimentos de obter dados do sensor;
- implementação de Service para procedimento de obter Mac do usuário;
- implementação de service de operações de GPS;

## Execução

Desenvolvido em ambiente trivial de execução como Android Studio.