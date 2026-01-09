package routes

import (
	"file-upload-service/controllers"
	"file-upload-service/middleware"

	"github.com/gin-gonic/gin"
)

func FileRoutes(router *gin.RouterGroup) {
	file_routes := router.Group("/file")

	file_routes.GET("/all", middleware.AuthRequired, controllers.ListFiles)
	file_routes.POST("/upload", middleware.AuthRequired, controllers.UploadFile)
	file_routes.GET("/:filename", middleware.AuthRequired, controllers.GetFile)
	file_routes.DELETE("/:filename", middleware.AuthRequired, controllers.DeleteFile)
}
