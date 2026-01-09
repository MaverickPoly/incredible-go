package main

import (
	"file-upload-service/routes"
	"file-upload-service/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Env Loading")
	godotenv.Load()
	utils.ConnectDB()
	fmt.Println("Env Loaded")
	app := gin.Default()

	routes.SetupRoutes(app)

	PORT := os.Getenv("PORT")

	fmt.Println("PORT:", PORT)
	app.Run(fmt.Sprintf(":%s", PORT))
}
