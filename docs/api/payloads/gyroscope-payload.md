### 📡 Endpoint: `/telemetry/gyroscope`

> This endpoint receives **gyroscope sensor data**, including axis values, timestamp, and the device's MAC address.

**Method:** `POST`
**Path:** `/telemetry/gyroscope`

#### 📥 Request Body (JSON)

```json
{
  "mac_address": "AB:CD:EF:12:34:56",
  "timestamp": "2025-06-16T10:00:00Z",
  "x": -0.012,
  "y": 0.987,
  "z": -9.802
}
```

#### 📄 Expected Fields

| Field         | Type   | Description                                 |
| ------------- | ------ | ------------------------------------------- |
| `mac_address` | string | Unique MAC address of the device            |
| `timestamp`   | string | ISO 8601 formatted timestamp of the reading |
| `x`           | float  | Gyroscope X-axis value                      |
| `y`           | float  | Gyroscope Y-axis value                      |
| `z`           | float  | Gyroscope Z-axis value                      |

#### ✅ Example Response (200 OK)

```json
{
  "message": "Gyroscope data received successfully"
}
```

#### ❌ Example Response (400 Bad Request)

Returned when the request body is missing required fields or contains invalid values.

```json
{
  "error": "Validation failed",
  "details": [
    {
      "field": "mac_address",
      "message": "mac_address is required"
    },
    {
      "field": "timestamp",
      "message": "timestamp must be a valid ISO 8601 string"
    },
    {
      "field": "x",
      "message": "x must be a numeric value"
    }
  ]
}
```

#### 🔎 Explanation

| Field         | Cause                                            |
| ------------- | ------------------------------------------------ |
| `mac_address` | Missing or null                                  |
| `timestamp`   | Invalid or non-ISO 8601 format (e.g., malformed) |
| `x`           | Not a number (e.g., string, null, or malformed)  |
