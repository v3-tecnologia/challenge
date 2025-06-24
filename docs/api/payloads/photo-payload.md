### 📸 Endpoint: `/telemetry/photo`

> This endpoint receives a **photo in Base64 format**, along with the timestamp and the device's MAC address.

**Method:** `POST`
**Path:** `/telemetry/photo`

#### 📥 Request Body (JSON)

```json
{
  "mac_address": "AB:CD:EF:12:34:56",
  "timestamp": "2025-06-16T10:00:00Z",
  "photo_base64": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD..."
}
```

#### 📄 Expected Fields

| Field          | Type   | Description                                       |
| -------------- | ------ | ------------------------------------------------- |
| `mac_address`  | string | Unique MAC address of the device                  |
| `timestamp`    | string | ISO 8601 formatted timestamp of the photo capture |
| `photo_base64` | string | Base64-encoded JPEG or PNG image.                 |

#### 📌 Notes

- The `photo_base64` must include the proper **MIME type prefix**, such as:
  - `"data:image/jpeg;base64,..."`
  - `"data:image/png;base64,..."`

#### ✅ Example Response (200 OK)

```json
{
  "message": "Photo received successfully"
}
```

---

#### ❌ Example Response (400 Bad Request)

Returned when the request body is invalid or missing required fields.

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
      "field": "photo_base64",
      "message": "photo_base64 must be a valid base64 JPEG or PNG image"
    }
  ]
}
```
