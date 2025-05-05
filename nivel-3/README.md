# Desafio V3 - Nível 1

No terceiro nível optei por alterar o script do primeiro nível em shell.

1. Transformar o arquivo em executável

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