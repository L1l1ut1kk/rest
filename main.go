package main

import (
	"github.com/L1l1ut1kk/rest/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// Connect to the database
	db, err := gorm.Open("sqlite3", "images.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	// Create the Image model in the database
	db.AutoMigrate(&models.Image{})

	// Route handler for image upload
	r.POST("/upload", func(c *gin.Context) {
		// Get the uploaded file
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(400, gin.H{"error": "img upload failed"})
			return
		}

		// Read the file data into a byte slice
		fileData, err := file.Open()
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to read img file"})
			return
		}
		defer fileData.Close()

		data := make([]byte, file.Size)
		_, err = fileData.Read(data)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to read img file"})
			return
		}

		// Create a new Image record in the database
		image := models.Image{Name: file.Filename, Data: data}
		db.Create(&image)

		c.JSON(200, gin.H{"status": "god damn, u did it"})
	})

	r.Run(":8080")
}
