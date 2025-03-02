package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/yashpatil74/nimbus/internal/domain/entities"
	"github.com/yashpatil74/nimbus/internal/repository"
)

type FileService struct {
	repo            *repository.FileRepository
	folderRepo      *repository.FolderRepository
	baseStoragePath string
}

func NewFileService(repo *repository.FileRepository, folderRepo *repository.FolderRepository, basePath string) *FileService {
	return &FileService{
		repo:            repo,
		folderRepo:      folderRepo,
		baseStoragePath: basePath,
	}
}

func (fs *FileService) UploadFile(userID, folderID string, file *multipart.FileHeader) (*entities.File, error) {
	if folderID != "" {
		folder, err := fs.folderRepo.GetFolderByID(folderID)
		if err != nil {
			return nil, err
		}
		if folder.UserID != userID {
			return nil, fmt.Errorf("unauthorized access to folder")
		}
	}

	fileID := uuid.New().String()

	fileEntity := &entities.File{
		ID:          fileID,
		Name:        file.Filename,
		ContentType: file.Header.Get("Content-Type"),
		Size:        file.Size,
		UserID:      userID,
		FolderID:    folderID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsEncrypted: true,
	}

	fileEntity.DetectType()

	encSuffix, err := generateRandomString(8)
	if err != nil {
		return nil, fmt.Errorf("failed to generate secure filename: %w", err)
	}

	fileExt := filepath.Ext(fileEntity.Name)
	fileEntity.EncryptedName = fmt.Sprintf("%s_%s%s", fileID, encSuffix, fileExt)

	userDir := filepath.Join(fs.baseStoragePath, userID)
	if err := os.MkdirAll(userDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create user directory: %w", err)
	}

	fileEntity.StoragePath = filepath.Join(userDir, fileEntity.EncryptedName)

	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(fileEntity.StoragePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	//Encryption (yet to implement)
	key, iv, err := generateEncryptionKeyAndIV()
	if err != nil {
		return nil, fmt.Errorf("failed to generate encryption key and IV: %w", err)
	}

	fileEntity.EncryptionKey = key
	fileEntity.EncryptionIV = iv

	if _, err := io.Copy(dst, src); err != nil {
		os.Remove(fileEntity.StoragePath)
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	if err := fs.repo.SaveFile(fileEntity); err != nil {
		os.Remove(fileEntity.StoragePath)

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, fmt.Errorf("a file with this name already exists in this folder")
		}
		return nil, fmt.Errorf("failed to save file metadata: %w", err)
	}

	return fileEntity, nil
}

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2) // hex encoding doubles the length
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func generateEncryptionKeyAndIV() (string, string, error) {
	// Generate 32-byte key (for AES-256)
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", "", err
	}

	// Generate 16-byte IV
	iv := make([]byte, 16)
	if _, err := rand.Read(iv); err != nil {
		return "", "", err
	}

	return hex.EncodeToString(key), hex.EncodeToString(iv), nil
}
