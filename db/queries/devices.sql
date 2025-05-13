-- name: InsertDevice :one
INSERT INTO devices (device_id) VALUES ($1) RETURNING *;

-- name: GetDeviceByID :one
SELECT * FROM devices WHERE device_id = $1;
