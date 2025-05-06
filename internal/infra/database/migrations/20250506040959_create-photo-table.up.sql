CREATE TABLE IF NOT EXISTS photo (
  id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  file_path TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  deleted_at TIMESTAMP
);