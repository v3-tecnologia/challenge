-- name: InsertGyroscopeReading :one
INSERT INTO gyroscope_readings (device_id, x, y, z, collected_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

