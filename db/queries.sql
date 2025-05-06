-- name: InsertDevice :one
INSERT INTO devices (device_id, name, model) VALUES ($1, $2, $3) RETURNING *;

