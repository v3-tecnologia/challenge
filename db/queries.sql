-- name: InsertDevice :one
INSERT INTO devices (device_id, name, model) VALUES ($1, $2, $3) RETURNING *;

-- name: InsertGyroscopeReading :one
INSERT INTO gyroscope_readings (device_id, x, y, z, collected_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

