ALTER TABLE files ADD COLUMN encrypted_name TEXT;
ALTER TABLE files ADD COLUMN encryption_key TEXT;
ALTER TABLE files ADD COLUMN encryption_iv TEXT;
ALTER TABLE files ADD COLUMN is_encrypted BOOLEAN NOT NULL DEFAULT 0;

CREATE INDEX idx_files_is_encrypted ON files(is_encrypted);