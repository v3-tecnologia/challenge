Ajustador de Volume para Veículos

Solução em Java para controle de volume de veículos via linha de comando ou arquivo CSV com validações inteligentes.

Como Usar

Modo Individual
Ajuste o volume de um único veículo:

java -jar ajustar_volume.jar --placa ABC1D23 --volume 50
Modo em Lote (CSV)
Processe múltiplos veículos de um arquivo CSV:

java -jar ajustar_volume.jar --arquivo veiculos.csv
Arquivo CSV de Exemplo (veiculos.csv)

placa,volume
ABC1D23,50
DEF4E56,30
GHI7F89,40
JKL1M23,15
Regras do CSV:

Primeira linha deve ser cabeçalho: placa,volume
Uma placa por linha
Formato de placa válido: ABC1D23 (3 letras + 1 número + 1 letra + 2 números)
Volume deve ser entre 0 e 100
Funcionalidades

Validação automática de placas
Checagem de volume (0-100)
Processamento paralelo para arquivos grandes
Logs detalhados em tempo real
Mensagens de Erro Comuns

Erro	Causa	Solução
"Formato de placa inválido!"	Placa não segue padrão BR	Use formato AAA1A11
"Volume deve ser 0-100!"	Valor fora do range	Ajuste para valor válido
"Arquivo não encontrado!"	Caminho incorreto	Verifique local do arquivo
Compilação e Execução

Compile:

javac src/*.java -d build
Execute:

java -cp build Main --placa ABC1D23 --volume 50
Estrutura do Projeto

src/
├── Main.java              (Ponto de entrada)
├── Vehicle.java           (Modelo de veículo)
├── VolumeAdjuster.java    (Núcleo do ajuste)
├── VehicleProcessor.java  (Processamento em lote)
└── CustomLogger.java      (Sistema de logs)
Dicas

Para testar rapidamente:

echo "placa,volume\nABC1D23,50" > teste.csv && java -jar ajustar_volume.jar --arquivo teste.csv
Visualize os logs em tempo real com:

java -jar ajustar_volume.jar --arquivo veiculos.csv | grep "Ajustando"
