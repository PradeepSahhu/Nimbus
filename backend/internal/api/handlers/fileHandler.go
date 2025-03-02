package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yashpatil74/nimbus/internal/api/middlewares"
	"github.com/yashpatil74/nimbus/internal/services"
)

type FileHandler struct {
	fileService *services.FileService
}

func NewFileHandler(FileService *services.FileService) *FileHandler {
	return &FileHandler{
		fileService: FileService,
	}
}

func (fh *FileHandler) UploadFile(c *gin.Context) {
	userID, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	folderID := c.DefaultQuery("folderId", "")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	uploadedFile, err := fh.fileService.UploadFile(userID, folderID, file)
	if err != nil {
		// Determine appropriate error status code based on error type
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "folder not found") {
			statusCode = http.StatusNotFound
		} else if strings.Contains(err.Error(), "unauthorized") {
			statusCode = http.StatusForbidden
		} else if strings.Contains(err.Error(), "already exists") {
			statusCode = http.StatusConflict
		}

		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "File uploaded successfully",
		"file":    uploadedFile,
	})
}
