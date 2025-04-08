# Desafio Suporte Técnico

## 💻 O Problema

### Contexto

Dispositivos embarcados quase sempre podem receber comandos de forma remota. Isso é possível através de uma comunicação bidirecional entre o servidor central e os dispositivos instalados nos veículos. Por exemplo, quando um dispositivo detecta uma frenagem brusca, ele envia essa informação para o servidor. Da mesma forma, o servidor pode enviar comandos para o dispositivo, como ajustar configurações, atualizar firmware ou, como neste caso, alterar o volume dos alertas sonoros.

### Problema Específico

Os motoristas da empresa **Transportes Rápidos Brasil** estão reclamando que os sinais sonoros dos dispositivos embarcados estão muito altos. Esses sinais são ativados quando o dispositivo detecta comportamentos específicos do condutor, como curvas bruscas ou frenagens intensas.

A empresa solicitou com urgência que o volume desses alertas seja reduzido em uma lista específica de veículos.

### Solução Necessária

Existe um endpoint que permite enviar comandos diretamente para os dispositivos. Este endpoint aceita requisições HTTP com os seguintes parâmetros:

- `PLACA`: Identificação do veículo
- `VOLUME_ALERTA`: Valor do volume (0-100)

**Exemplo de Requisição HTTP:**
```
POST /api/dispositivos/configurar
Content-Type: application/json

{
  "placa": "ABC1234",
  "volume_alerta": 50
}
```

## 🚀 Bora nessa!

Você deverá criar uma ferramenta CLI em Python para enviar comandos de ajuste de volume para os dispositivos embarcados.

## Níveis

### Nível 1

Crie um programa que receba por parâmetro uma placa e um volume, e envie o comando para o dispositivo correspondente.

**Exemplo de uso:**
```
python ajustar_volume.py --placa ABC1234 --volume 50
```

### Nível 2

Modifique o programa para receber como parâmetro um arquivo CSV com todas as placas e os volumes correspondentes, e execute os comandos em batch.

**Exemplo de arquivo CSV:**
```
placa,volume
ABC1234,50
DEF5678,30
GHI9012,40
```

**Exemplo de uso:**
```
python ajustar_volume.py --arquivo veiculos.csv
```

### Nível 3

Implemente a execução paralela dos comandos para processar múltiplas requisições simultaneamente, tornando o processo mais eficiente.

**Dica:** Utilize bibliotecas como `concurrent.futures` ou `asyncio` para implementar a paralelização.

### Nível 4

Adicione validação de dados e tratamento de erros ao programa:
- Verifique se a placa está em um formato válido
- Verifique se o volume está dentro do intervalo permitido (0-100)
- Trate erros de conexão com o servidor
- Implemente logs para acompanhar o progresso das operações

### Nível 5

Crie testes unitários para garantir que o programa funcione corretamente em diferentes cenários.

## 📝 Dicas

- Utilize bibliotecas como `argparse` ou `click` para criar interfaces de linha de comando
- Considere usar `pandas` para manipulação de arquivos CSV
- Implemente um sistema de retry para lidar com falhas temporárias de conexão
- Documente bem o código e crie um README explicando como usar a ferramenta

## ⏰ Tempo para Entrega

Quanto antes você enviar, mais cuidado podemos ter na revisão do seu teste. Faça no seu tempo, mas mantenha a qualidade!

**Mas não desista! Envie até onde conseguir.**

Boa sorte! 🍀
