CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS devices (
  device_id TEXT PRIMARY KEY,
  registered_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS gyroscope_readings (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  device_id TEXT NOT NULL REFERENCES devices(device_id) ON DELETE CASCADE,
  x FLOAT NOT NULL,
  y FLOAT NOT NULL,
  z FLOAT NOT NULL,
  collected_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS gps_readings (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  device_id TEXT NOT NULL REFERENCES devices(device_id) ON DELETE CASCADE,
  latitude DECIMAL(9,6) NOT NULL,
  longitude DECIMAL(9,6) NOT NULL,
  collected_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS photos (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  device_id TEXT NOT NULL REFERENCES devices(device_id) ON DELETE CASCADE,
  image_url TEXT NOT NULL,
  collected_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_gyro_device_time ON gyroscope_readings(device_id, collected_at);
CREATE INDEX idx_gps_device_time ON gps_readings(device_id, collected_at);
CREATE INDEX idx_photo_device_time ON photos(device_id, collected_at);
