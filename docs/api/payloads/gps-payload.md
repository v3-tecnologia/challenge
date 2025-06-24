### 🧭 Endpoint: `/telemetry/gps`

> This endpoint receives **GPS coordinates**, including latitude, longitude, timestamp, and the device's MAC address.

**Method:** `POST`
**Path:** `/telemetry/gps`

#### 📥 Request Body (JSON)

```json
{
  "mac_address": "AB:CD:EF:12:34:56",
  "timestamp": "2025-06-16T10:00:00Z",
  "latitude": -28.481673,
  "longitude": -49.013159
}
```

#### 📄 Expected Fields

| Field         | Type   | Description                                 |
| ------------- | ------ | ------------------------------------------- |
| `mac_address` | string | Unique MAC address of the device            |
| `timestamp`   | string | ISO 8601 formatted timestamp of the reading |
| `latitude`    | number | Latitude coordinate (decimal degrees)       |
| `longitude`   | number | Longitude coordinate (decimal degrees)      |

#### ✅ Example Response (200 OK)

```json
{
  "message": "GPS data received successfully"
}
```

---

#### ❌ Example Response (400 Bad Request)

Returned when required fields are missing or contain invalid values.

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
      "field": "latitude",
      "message": "latitude must be a valid number"
    },
    {
      "field": "longitude",
      "message": "longitude must be a valid number"
    }
  ]
}
```
