package controllers

import "github.com/gin-gonic/gin"

type RouterController struct {
}

func NewRouterController() *RouterController {
	return &RouterController{}
}

func (rc *RouterController) CheckHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}
