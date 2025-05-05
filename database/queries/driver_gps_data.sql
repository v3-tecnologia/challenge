-- name: SaveGPS :exec
-- Insert a new GPS record
-- Parameters: id, device_id, latitude, longitude, timestamp, created_at
INSERT INTO driver_gps_data (
    id,
    device_id,
    latitude,
    longitude,
    timestamp,
    created_at
)
VALUES ($1, $2, $3, $4, $5, $6);
