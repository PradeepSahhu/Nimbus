package entities

import (
	"path/filepath"
	"strings"
	"time"
)

type File struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	EncryptedName string    `json:"-"`
	EncryptionKey string    `json:"-"`
	EncryptionIV  string    `json:"-"`
	ContentType   string    `json:"content_type"`
	Type          string    `json:"type"`
	Size          int64     `json:"size"`
	UserID        string    `json:"user_id"`
	FolderID      string    `json:"folder_id"`
	StoragePath   string    `json:"-"`
	IsEncrypted   bool      `json:"is_encrypted"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (f *File) GetExtension() string {
	return strings.TrimPrefix(filepath.Ext(f.Name), ".")
}

func (f *File) InRootFolder() bool {
	return f.FolderID == ""
}

func (f *File) DetectType() {
	contentType := strings.ToLower(f.ContentType)

	switch {
	case strings.Contains(contentType, "image/"):
		f.Type = "image"
	case strings.Contains(contentType, "video/"):
		f.Type = "video"
	case strings.Contains(contentType, "audio/"):
		f.Type = "audio"
	case strings.Contains(contentType, "text/") ||
		strings.Contains(contentType, "application/pdf") ||
		strings.Contains(contentType, "application/msword") ||
		strings.Contains(contentType, "application/vnd.openxmlformats"):
		f.Type = "document"
	default:
		f.Type = "other"
	}
}
