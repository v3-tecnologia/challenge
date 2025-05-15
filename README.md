# Go Telemetry API

This is a Go application that provides an API for receiving telemetry data from a device. It has three endpoints for gyroscope data, GPS data, and photo data.

### Requirements/Setup

1. GO Version 1.24.3
Download and install Go following the steps in the official documentation: [Go Installation Guide](https://go.dev/doc/install#install).

2. Docker (28.1.1+) and docker-compose (v2.6.0+)
Install Docker engine https://docs.docker.com/engine/install/


### Setup

##### Clone your repo (if needed)
```
git clone [<repo-url>](https://github.com/martinsrenan/challenge)
```

```
cd challenge
```

##### Start services
```
docker-compose up -d
```

This will start containers for the app and postgres database.

### Test endpoints

Test endpoints using following curl commands via terminal:

#### Endpoint: /telemetry/gyroscope
```
curl -X POST http://localhost:8080/telemetry/gyroscope \
-H "Content-Type: application/json" \
-d '{"x": 1.0, "y": 2.0, "z": 3.0}'
```
Expected success Message:
{"message":"Gyroscope data received successfully"}


##### Error Cenario:
```
curl -X POST http://localhost:8080/telemetry/gyroscope \
-H "Content-Type: application/json" \
-d '{"x": 0, "y": 0, "z": 0}'
```

Expected error Message:
"Missing or invalid gyroscope data"


#### Endpoint: /telemetry/gps
```
curl -X POST http://localhost:8080/telemetry/gps \
-H "Content-Type: application/json" \
-d '{"latitude": 37.7749, "longitude": -122.4194, "altitude": 30.0}'
```

Expected success Message:
{"message":"Gyroscope data received successfully"}


##### Error Cenario:
```
curl -X POST http://localhost:8080/telemetry/gps \
-H "Content-Type: application/json" \
-d '{"latitude": 0, "longitude": 0, "altitude": 30.0}'
```
Expected error Message:
"Missing or invalid GPS data"


#### Endpoint: /telemetry/photo
```
curl -X POST http://localhost:8080/telemetry/photo \
-H "Content-Type: application/json" \
-d '{"filename": "image.png", "data": "base64encodeddata"}'
```

Expected success Message:
{"message":"Photo data received successfully"}


##### Error Cenario:
```
curl -X POST http://localhost:8080/telemetry/photo \
-H "Content-Type: application/json" \
-d '{"filename": "", "data": ""}'
```

Expected error Message:
"Missing photo data fields"


# TO DO List

### Nível 3

Crie testes unitários para cada arquivo da aplicação. Para cada nova implementação a seguir, também deve-se criar os testes.
Improve doc with unity test information

### Nível 5

A cada foto recebida, deve-se utilizar o AWS Rekognition para comparar se a foto enviada é reconhecida com base nas fotos anteriores enviadas.

Se a foto enviada for reconhecida, retorne como resposta do `POST` um atributo que indique isso.

Utilize as fotos iniciais para realizar o treinamento da IA.

