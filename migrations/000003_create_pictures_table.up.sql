CREATE TABLE IF NOT EXISTS pictures (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    device_id VARCHAR(255) NOT NULL,
    received_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    picture_url VARCHAR(512) NOT NULL,
    picture_format VARCHAR(50) NOT NULL,
    recognized_face BOOLEAN DEFAULT FALSE,
    rekognition_score DOUBLE PRECISION
    );

CREATE INDEX IF NOT EXISTS idx_photo_device_id ON pictures(device_id);
CREATE INDEX IF NOT EXISTS idx_photo_created_at ON pictures(created_at DESC);
