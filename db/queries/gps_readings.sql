-- name: InsertGPSReading :one
INSERT INTO
    gps_readings (device_id, latitude, longitude, collected_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

