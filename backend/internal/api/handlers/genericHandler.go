package handlers

import "github.com/gin-gonic/gin"

type GenericHandler struct {
}

func NewGenericHandler() *GenericHandler {
	return &GenericHandler{}
}

func (h *GenericHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
