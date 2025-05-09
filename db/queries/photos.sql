-- name: InsertPhoto :one
INSERT INTO photos (device_id, image_url, collected_at) VALUES ($1, $2, $3) RETURNING *;

