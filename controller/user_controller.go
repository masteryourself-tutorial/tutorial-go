package controller

import (
	"chat/service"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags 获取用户列表
// @Success 200
// @Router /user/list [get]
func GetUserList(ctx *gin.Context) {
	data := service.GetUserList()
	ctx.JSON(200, gin.H{
		"message": data,
	})
}
