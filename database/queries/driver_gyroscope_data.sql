-- name: SaveGyroscope :exec
-- Insert a new gyroscope reading
-- Parameters: id, device_id, x, y, z, timestamp, created_at
INSERT INTO driver_gyroscope_data (
    id,
    device_id,
    x,
    y,
    z,
    timestamp,
    created_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7);
