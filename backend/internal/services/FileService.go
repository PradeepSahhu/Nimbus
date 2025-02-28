package services

import "github.com/yashpatil74/nimbus/internal/repository"

type FileService struct {
	repo *repository.FileRepository
}

func NewFileService(repo *repository.FileRepository) *FileService {
	return &FileService{
		repo: repo,
	}
}
