POST http://localhost:3000/telemetry/gyroscope
{
  "unique_id": "1",
  "timestamp": "2",
  "x": 3.0,
  "y": 4.0,
  "z": 5.0
}
HTTP 200
Content-Type: application/json
[Asserts]
jsonpath "$.status" == "ok"