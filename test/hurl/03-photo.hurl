POST http://localhost:3000/telemetry/photo
{
  "unique_id": "1",
  "timestamp": "2",
  "photo": "data:image/png;base64"
}
HTTP 200
Content-Type: application/json
[Asserts]
jsonpath "$.status" == "ok"
