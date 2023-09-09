package router

import (
	"chat/controller"
	_ "chat/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	engine := gin.Default()
	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// user
	engine.
		GET("/user", controller.SelectUser).
		POST("/user", controller.CreateUser).
		PUT("/user", controller.UpdateUser).
		DELETE("/user", controller.DeleteUser).
		GET("/user/list", controller.UserList)
	return engine
}
