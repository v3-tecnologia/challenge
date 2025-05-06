-- name: SavePhoto :exec
-- Insert a new photo record
-- Parameters: id, device_id, file_path, recognized, timestamp, created_at
INSERT INTO driver_photos (
    id,
    device_id,
    file_path,
    recognized,
    timestamp,
    created_at
)
VALUES ($1, $2, $3, $4, $5, $6);
