DROP INDEX IF EXISTS idx_files_is_encrypted;

ALTER TABLE files DROP COLUMN encrypted_name;
ALTER TABLE files DROP COLUMN encryption_key;
ALTER TABLE files DROP COLUMN encryption_iv;
ALTER TABLE files DROP COLUMN is_encrypted;