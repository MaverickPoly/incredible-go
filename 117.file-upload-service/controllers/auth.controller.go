package controllers

import (
	"file-upload-service/models"
	"file-upload-service/utils"
	"fmt"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid body!"})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "Some fields are missing!"})
		return
	}

	var dbUser models.User
	if err := utils.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("User with username %s not found!", user.Username)})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid password!"})
		return
	}

	accessToken, _ := utils.GenerateToken(dbUser)
	c.SetCookie(
		"accessToken",
		accessToken,
		60*60*24*7,
		"/",
		"localhost",
		false,
		false,
	)

	c.JSON(200, gin.H{"accessToken": accessToken, "message": "Logged in successfully!"})
}

func HandleRegister(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid body!"})
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "Some fields are missing!"})
		return
	}

	var existingUser models.User
	if err := utils.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{"error": "Username or email already exists!"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error hashing password: %s", err.Error())})
		return
	}

	user.Password = string(hashedPassword)

	if err := utils.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error creating user: %s", err.Error())})
		return
	}

	c.JSON(201, user)
}

func HandleLogout(c *gin.Context) {
	c.SetCookie(
		"accessToken",
		"",
		-1,
		"/",
		"localhost",
		false,
		false,
	)

	c.JSON(200, gin.H{"message": "Logged out successfully!"})
}

// Get Current User
func GetUser(c *gin.Context) {
	userId, exists := c.Get("userId")

	if !exists {
		c.JSON(401, gin.H{"error": "Could not get user id!"})
		return
	}

	var user models.User
	if err := utils.DB.Find(&user, userId).Error; err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error fetching profile: %s", err.Error())})
		return
	}

	if user.ID == 0 {
		c.JSON(500, gin.H{"error": "Could not find user!"})
		return
	}

	c.JSON(200, user)
}
