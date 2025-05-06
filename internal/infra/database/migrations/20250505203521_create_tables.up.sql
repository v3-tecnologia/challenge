CREATE TABLE IF NOT EXISTS gyroscope (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    x_value NUMERIC NOT NULL,
    y_value NUMERIC NOT NULL,
    z_value NUMERIC NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS geolocation (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    latitude NUMERIC NOT NULL,
    longitude NUMERIC NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);