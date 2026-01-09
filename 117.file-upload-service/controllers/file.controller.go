package controllers

import (
	"file-upload-service/models"
	"file-upload-service/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const uploadDir = "./uploads"

func UploadFile(c *gin.Context) {
	userId, exists := c.Get("userId")

	if !exists {
		c.JSON(401, gin.H{"error": "Could not get user id!"})
		return
	}

	var user models.User
	if err := utils.DB.Find(&user, userId).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error fetching user!"})
		return
	}

	// Retrieve the file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No file provided!"})
		return
	}

	// Validate file size
	file_size_mb := file.Size / (1024 * 1024)
	if file_size_mb > 100 {
		c.JSON(400, gin.H{"error": "File too large. Limit is 100MB"})
		return
	}

	filename := filepath.Base(file.Filename)
	userDir := filepath.Join(uploadDir, user.Username)
	if err := os.MkdirAll(userDir, 0755); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user directory!"})
		return
	}

	dst := filepath.Join(uploadDir, user.Username, filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save file!"})
		return
	}

	fileModel := models.File{
		UserId:   user.ID,
		Path:     dst,
		Filename: filename,
		Size:     file.Size,
	}

	if err := utils.DB.Create(&fileModel).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error creating file!"})
		return
	}

	c.JSON(201, gin.H{"message": "File uploaded successfully!", "url": fmt.Sprintf("/file/%s", file.Filename)})
}

func GetFile(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"error": "Could not get user id!"})
		return
	}

	var user models.User
	if err := utils.DB.Find(&user, userId).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error fetching user!"})
		return
	}

	filename := c.Param("filename")
	filePath := filepath.Join(uploadDir, user.Username, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(404, gin.H{"error": "File not found!"})
		return
	}

	c.File(filePath)
}

func ListFiles(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"error": "Could not get user id!"})
		return
	}

	var user models.User
	if err := utils.DB.Find(&user, userId).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error fetching user!"})
		return
	}

	var files []models.File
	if err := utils.DB.Where("user_id = ?", user.ID).Find(&files).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error fetching files!"})
		return
	}

	c.JSON(200, files)
}

func DeleteFile(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"error": "Could not get user id!"})
		return
	}

	var user models.User
	if err := utils.DB.Find(&user, userId).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error fetching user!"})
		return
	}

	filename := c.Param("filename")
	if filename == "" {
		c.JSON(400, gin.H{"error": "Filename is required!"})
		return
	}

	safeFilename := filepath.Base(filename)
	var fileRecord models.File
	if err := utils.DB.Where("user_id = ? AND filename = ?", user.ID, safeFilename).First(&fileRecord).Error; err != nil {
		c.JSON(404, gin.H{"error": "File not found!"})
		return
	}

	// Delete it from filesystem
	filePath := filepath.Join(uploadDir, user.Username, safeFilename)
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		c.JSON(500, gin.H{"error": "Failed to delete file from disk!"})
		return
	}

	// Delete it from database
	if err := utils.DB.Delete(&fileRecord).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete file record!"})
		return
	}

	c.JSON(200, gin.H{"message": "File deleted successfully!"})
}
