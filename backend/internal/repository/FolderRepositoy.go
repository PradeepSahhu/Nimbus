package repository

import (
	"database/sql"
	"fmt"

	"github.com/yashpatil74/nimbus/internal/domain/entities"
)

type FolderRepository struct {
	db *sql.DB
}

func NewFolderRepository(db *sql.DB) *FolderRepository {
	return &FolderRepository{
		db: db,
	}
}

func (fr *FolderRepository) GetFolderByID(folderID string) (*entities.Folder, error) {
	query := `
        SELECT id, name, user_id, parent_id, created_at, updated_at
        FROM folders
        WHERE id = ?
    `

	folder := &entities.Folder{}

	var parentID sql.NullString // Handle NULL parent_id values

	err := fr.db.QueryRow(query, folderID).Scan(
		&folder.ID,
		&folder.Name,
		&folder.UserID,
		&parentID,
		&folder.CreatedAt,
		&folder.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("folder not found with ID: %s", folderID)
		}
		return nil, fmt.Errorf("database error when fetching folder: %w", err)
	}

	if parentID.Valid {
		folder.ParentID = parentID.String
	} else {
		folder.ParentID = ""
	}

	return folder, nil
}
