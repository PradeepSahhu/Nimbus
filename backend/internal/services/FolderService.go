package services

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/yashpatil74/nimbus/internal/repository"
)

type FolderService struct {
	folderRepo      *repository.FolderRepository
	baseStoragePath string
}

func NewFolderService(folderRepo *repository.FolderRepository, basePath string) *FolderService {
	return &FolderService{
		folderRepo:      folderRepo,
		baseStoragePath: basePath,
	}
}

func (fs *FolderService) CreateUserFolder(userId string) error {
	if userId == "" {
		return errors.New("user ID cannot be empty")
	}

	userStoragePath := filepath.Join(fs.baseStoragePath, userId)
	if err := os.MkdirAll(userStoragePath, 0755); err != nil {
		return err
	}

	return nil
}
