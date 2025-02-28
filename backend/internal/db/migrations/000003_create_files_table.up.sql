CREATE TABLE IF NOT EXISTS files (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    content_type VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    size BIGINT NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    folder_id VARCHAR(36),
    storage_path VARCHAR(512) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (folder_id) REFERENCES folders(id) ON DELETE SET NULL
);

CREATE INDEX idx_files_user_id ON files(user_id);
CREATE INDEX idx_files_folder_id ON files(folder_id);
CREATE INDEX idx_files_type ON files(type);

CREATE UNIQUE INDEX idx_unique_file_name_per_folder ON files(user_id, folder_id, name);

CREATE INDEX idx_files_name ON files(name);