package handlers

import (
	"github.com/gin-gonic/gin"
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

}
