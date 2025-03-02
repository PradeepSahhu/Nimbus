package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/yashpatil74/nimbus/internal/domain/entities"
)

type FileRepository struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

func (fr *FileRepository) SaveFile(file *entities.File) error {
	query := `
        INSERT INTO files (
            id, name, encrypted_name, content_type, type, size, 
            user_id, folder_id, storage_path, encryption_key,
            encryption_iv, is_encrypted, created_at, updated_at
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	_, err := fr.db.Exec(
		query,
		file.ID,
		file.Name,
		file.EncryptedName,
		file.ContentType,
		file.Type,
		file.Size,
		file.UserID,
		file.FolderID,
		file.StoragePath,
		file.EncryptionKey,
		file.EncryptionIV,
		file.IsEncrypted,
		file.CreatedAt,
		file.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to insert file record: %w", err)
	}

	return nil
}

func (fr *FileRepository) DeleteFile(fileID string, userID string) error {
	query := `DELETE FROM files WHERE id = ? AND user_id = ?`

	result, err := fr.db.Exec(query, fileID, userID)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("file not found or unauthorized access")
	}

	return nil
}

func (fr *FileRepository) GetFileByID(fileID string, userID string) (*entities.File, error) {
	query := `
        SELECT id, name, encrypted_name, content_type, type, size,
               user_id, folder_id, storage_path, encryption_key,
               encryption_iv, is_encrypted, created_at, updated_at
        FROM files
        WHERE id = ? AND user_id = ?
    `

	file := &entities.File{}
	err := fr.db.QueryRow(query, fileID, userID).Scan(
		&file.ID,
		&file.Name,
		&file.EncryptedName,
		&file.ContentType,
		&file.Type,
		&file.Size,
		&file.UserID,
		&file.FolderID,
		&file.StoragePath,
		&file.EncryptionKey,
		&file.EncryptionIV,
		&file.IsEncrypted,
		&file.CreatedAt,
		&file.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("file not found or unauthorized access")
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	return file, nil
}

func (fr *FileRepository) GetFilesByUserID(userID string) ([]*entities.File, error) {
	query := `
		SELECT id, name, encrypted_name, content_type, type, size,
			   user_id, folder_id, storage_path, encryption_key,
			   encryption_iv, is_encrypted, created_at, updated_at
		FROM files
		WHERE user_id = ?
		ORDER BY name ASC
	`

	rows, err := fr.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query files: %w", err)
	}
	defer rows.Close()

	var files []*entities.File
	for rows.Next() {
		file := &entities.File{}
		err := rows.Scan(
			&file.ID,
			&file.Name,
			&file.EncryptedName,
			&file.ContentType,
			&file.Type,
			&file.Size,
			&file.UserID,
			&file.FolderID,
			&file.StoragePath,
			&file.EncryptionKey,
			&file.EncryptionIV,
			&file.IsEncrypted,
			&file.CreatedAt,
			&file.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan file row: %w", err)
		}
		files = append(files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return files, nil
}

func (fr *FileRepository) GetFilesByFolderID(folderID string, userID string) ([]*entities.File, error) {
	query := `
        SELECT id, name, content_type, type, size,
               user_id, folder_id, created_at, updated_at
        FROM files
        WHERE user_id = ? AND folder_id = ?
        ORDER BY name ASC
    `

	// For files in root folder
	if folderID == "" {
		query = `
            SELECT id, name, content_type, type, size,
                   user_id, folder_id, created_at, updated_at
            FROM files
            WHERE user_id = ? AND (folder_id = '' OR folder_id IS NULL)
            ORDER BY name ASC
        `
	}

	rows, err := fr.db.Query(query, userID, folderID)
	if err != nil {
		return nil, fmt.Errorf("failed to query files: %w", err)
	}
	defer rows.Close()

	var files []*entities.File
	for rows.Next() {
		file := &entities.File{}
		err := rows.Scan(
			&file.ID,
			&file.Name,
			&file.ContentType,
			&file.Type,
			&file.Size,
			&file.UserID,
			&file.FolderID,
			&file.CreatedAt,
			&file.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan file row: %w", err)
		}
		files = append(files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return files, nil
}
