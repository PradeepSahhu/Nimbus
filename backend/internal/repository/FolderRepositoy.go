package repository

import (
	"database/sql"
)

type FolderRepository struct {
	db *sql.DB
}

func NewFolderRepository(db *sql.DB) *FolderRepository {
	return &FolderRepository{
		db: db,
	}
}
