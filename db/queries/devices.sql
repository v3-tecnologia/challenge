-- name: InsertDevice :one
INSERT INTO devices (device_id, name, model) VALUES ($1, $2, $3) RETURNING *;

-- name: GetDeviceByID :one
SELECT * FROM devices WHERE device_id = $1;
