package handlers

import (
	"fmt"
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

func (fh *FileHandler) DownloadFile(c *gin.Context) {
	userID, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	fileID := c.Param("id")

	osFile, fileInfo, err := fh.fileService.DownloadFile(fileID, userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			statusCode = http.StatusNotFound
		} else if strings.Contains(err.Error(), "unauthorized") {
			statusCode = http.StatusForbidden
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	defer osFile.Close()

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name))
	c.Header("Content-Type", fileInfo.ContentType)

	c.File(osFile.Name())
}

func (fh *FileHandler) ListFiles(c *gin.Context) {
	userID, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	files, err := fh.fileService.ListFiles(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}

func (fh *FileHandler) DeleteFile(c *gin.Context) {
	userID, exists := middlewares.GetUserID(c)
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	fileID := c.Param("id")

	err := fh.fileService.DeleteFile(fileID, userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			statusCode = http.StatusNotFound
		} else if strings.Contains(err.Error(), "unauthorized") {
			statusCode = http.StatusForbidden
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
