CREATE TABLE IF NOT EXISTS gps_telemetry (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    device_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    received_at TIMESTAMP NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL
    );

CREATE INDEX IF NOT EXISTS idx_gps_device_id ON gps_telemetry(device_id);
CREATE INDEX IF NOT EXISTS idx_gps_timestamp ON gps_telemetry(created_at DESC);
