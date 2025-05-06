### 1. Enviar dados de GPS

- Envia latitude, longitude e timestamp do dispositivo.
- Exemplo de requisição

```bash
curl --location 'http://localhost:8080/telemetry/gps' \
--header 'Content-Type: application/json' \
--data-raw '{
  "device_id": "00:11:22:33:44",
  "latitude": -23.5505,
  "longitude": -46.6333,
  "timestamp": "2025-05-05T14:00:00Z"
}'
```

### 2. Enviar dados do Giroscópio

- Envia os dados x, y, z do giroscópio com timestamp.
- Exemplo de requisição

```bash
curl --location 'http://localhost:8080/telemetry/gyroscope' \
--header 'Content-Type: application/json' \
--data-raw '{
  "device_id": "00:11:22:33:44",
  "x": 0.12,
  "y": -0.03,
  "z": 0.87,
  "timestamp": "2025-05-05T14:00:00Z"
}'
```

### 3. Enviar foto

- Faz upload da imagem e envia timestamp. A API responde se o rosto foi reconhecido.
- Exemplo de requisição

```bash
curl --location 'http://localhost:8080/telemetry/photo' \
--form 'device_id="00:11:22:33:44"' \
--form 'timestamp="2025-05-05T14:00:00Z"' \
--form 'photo=@"/Users/mkafonso/Desktop/img.webp"'
```
