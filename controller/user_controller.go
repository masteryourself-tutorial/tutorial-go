package controller

import (
	"chat/model"
	"chat/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// SelectUser
// @Tags 用户模块
// @Summary 查询用户
// @Param id query string false "id"
// @Router /user/ [GET]
func SelectUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	data := service.SelectUser(uint(id))
	ctx.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Tags 用户模块
// @Summary 创建用户
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Router /user/ [POST]
func CreateUser(ctx *gin.Context) {
	name := ctx.Query("name")
	password := ctx.Query("password")
	user := model.User{
		Name:     name,
		Password: password,
	}
	service.CreateUser(user)
	ctx.JSON(200, gin.H{
		"message": "用户创建成功",
	})
}

// UpdateUser
// @Tags 用户模块
// @Summary 编辑用户
// @Param id query string false "id"
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Router /user/ [PUT]
func UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	name := ctx.Query("name")
	password := ctx.Query("password")
	user := model.User{
		Model: gorm.Model{
			ID: uint(id),
		},
		Name:     name,
		Password: password,
	}
	service.UpdateUser(user)
	ctx.JSON(200, gin.H{
		"message": "用户编辑成功",
	})
}

// DeleteUser
// @Tags 用户模块
// @Summary 删除用户
// @Param id query string false "id"
// @Router /user/ [DELETE]
func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	user := model.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	service.DeleteUser(user)
	ctx.JSON(200, gin.H{
		"message": "用户删除成功",
	})
}

// UserList
// @Tags 用户模块
// @Summary 获取用户列表
// @Router /user/list [GET]
func UserList(ctx *gin.Context) {
	data := service.GetUserList()
	ctx.JSON(200, gin.H{
		"message": data,
	})
}
