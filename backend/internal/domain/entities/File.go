package entities

import (
	"path/filepath"
	"strings"
	"time"
)

type File struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ContentType string    `json:"content_type"`
	Type        string    `json:"type"`
	Size        int64     `json:"size"`
	UserID      string    `json:"user_id"`
	FolderID    string    `json:"folder_id"`
	StoragePath string    `json:"storage_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (f *File) GetExtension() string {
	return strings.TrimPrefix(filepath.Ext(f.Name), ".")
}

func (f *File) InRootFolder() bool {
	return f.FolderID == ""
}

func DetectFileType(contentType string) string {
	contentType = strings.ToLower(contentType)

	switch {
	case strings.Contains(contentType, "image/"):
		return "image"
	case strings.Contains(contentType, "video/"):
		return "video"
	case strings.Contains(contentType, "audio/"):
		return "audio"
	case strings.Contains(contentType, "text/") ||
		strings.Contains(contentType, "application/pdf") ||
		strings.Contains(contentType, "application/msword") ||
		strings.Contains(contentType, "application/vnd.openxmlformats-officedocument"):
		return "document"
	default:
		return "other"
	}
}
