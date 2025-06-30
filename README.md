# V3 Challenge

Este projeto implementa um sistema de coleta, validação e persistência de dados de telemetria (giroscópio, GPS, fotos) utilizando:

- API HTTP em Go
- Comunicação assíncrona com NATS (Pub/Sub)
- Validação com JSON Schema
- Armazenamento em SQLite

## Diagrama de Alto Nível

![Diagrama](diag.png)

##  Endpoints HTTP

`
POST /telemetry/gyroscope (x, y, z)
`

`
POST /telemetry/gps       (latitude, longitude)
`

`
POST /telemetry/photo	  (png em base64)
`