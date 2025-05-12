CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS gyroscope_telemetry (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    device_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    received_at TIMESTAMP NOT NULL,
    x DOUBLE PRECISION NOT NULL,
    y DOUBLE PRECISION NOT NULL,
    z DOUBLE PRECISION NOT NULL
    );

CREATE INDEX IF NOT EXISTS idx_gyroscope_device_id ON gyroscope_telemetry(device_id);
CREATE INDEX IF NOT EXISTS idx_gyroscope_timestamp ON gyroscope_telemetry(created_at DESC);
