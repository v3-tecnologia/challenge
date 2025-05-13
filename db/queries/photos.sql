-- name: InsertPhoto :one
INSERT INTO
    photos (
    device_id,
    image_url,
    collected_at,
    recurrent_user
) VALUES ($1, $2, $3, $4) RETURNING *;

