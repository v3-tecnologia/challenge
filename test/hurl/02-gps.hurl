POST http://localhost:3000/telemetry/gps
{
  "unique_id": "1",
  "timestamp": "2",
  "latitude": 3.0,
  "longitude": 4.0
}
HTTP 200
Content-Type: application/json
[Asserts]
jsonpath "$.status" == "ok"