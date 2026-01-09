package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(app *gin.Engine) {
	router := app.Group("/")

	AuthRoutes(router)
	FileRoutes(router)
}
