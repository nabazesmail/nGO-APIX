package middleware

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Max size for the uploaded file (change it as needed)
		maxSize := int64(10 << 20) // 10 MB

		// Get the uploaded file from the form data
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from form data"})
			return
		}

		// Check file size
		if file.Size > maxSize {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the maximum allowed size"})
			return
		}

		// Create the uploads folder if it doesn't exist
		err = os.MkdirAll("assets/uploads", 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create uploads folder"})
			return
		}

		// Generate a unique filename for the uploaded file
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))

		// Save the file to the uploads folder
		err = c.SaveUploadedFile(file, filepath.Join("assets/uploads", filename))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
			return
		}

		// Store the file path in the context for later use in the controller
		c.Set("filePath", filepath.Join("assets/uploads", filename))

		c.Next()
	}
}
