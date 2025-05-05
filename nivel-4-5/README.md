# Desafio V3 - Nível 1

No primeiro nível optei por fazer script básico em shell.

Desafio:

Nível 4 e 5
Adicione validação de dados e tratamento de erros à sua solução:

* Verifique se a placa está em um formato válido
  - Usei regex para validar, segui o formato do arquivo de exemplo, se fosse real haveria a necessidade de outra regex.
* Verifique se o volume está dentro do intervalo permitido (0-100)
  - Usei as funções `lt` e `gt`
* Trate erros de conexão com o servidor
  - Dez tentativas com limite de 5 segundos de timeout.
* Implemente logs para acompanhar o progresso das operações
  - Escreve um arquivo de log.
* Transformar o arquivo em executável
  - Desde do primeiro nível já é um arquivo executável.
* Adicionar testes
  - Foram criados dois arquivos de exemplo. Um que passa durante a execução e outro que falhará. Isto valida o script!

1. Transforme o arquivo em executável:
```sh 
chmod +x ajustar_volume.sh
``` 

2. Crie um arquivo `.csv` com um conteúdo parecido com isto:

```csv
placa,volume_alerta
ABC1234,30
DEF5678,45
GHI9012,25
```

3. Executar o script passando o arquivo contendo os parâmetros e o endereço do servidor.

```sh
./ajustar_volume.sh -f veiculos.csv -u https://httpbin.org/post
``` 