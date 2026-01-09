package routes

import (
	"file-upload-service/controllers"
	"file-upload-service/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	auth_routes := router.Group("/auth")

	auth_routes.POST("/login", controllers.HandleLogin)
	auth_routes.POST("/register", controllers.HandleRegister)
	auth_routes.POST("/logout", middleware.AuthRequired, controllers.HandleLogout)
	auth_routes.GET("/me", middleware.AuthRequired, controllers.GetUser)
}
