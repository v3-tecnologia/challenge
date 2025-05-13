ALTER TABLE photos
ADD COLUMN recurrent_user BOOLEAN DEFAULT FALSE;

UPDATE photos
SET recurrent_user = FALSE
WHERE recurrent_user IS NULL;

