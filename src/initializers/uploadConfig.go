package initializers

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

// Helper function to check if the uploaded file is an image
func IsImageFile(fileHeader *multipart.FileHeader) bool {
	// Extract the file extension from the uploaded file's header
	ext := filepath.Ext(fileHeader.Filename)
	ext = strings.ToLower(ext)
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}

// Helper function to generate a unique filename for the uploaded image
func GenerateUniqueFilename(fileHeader *multipart.FileHeader) string {
	ext := filepath.Ext(fileHeader.Filename)
	// Create a unique filename using the original filename, timestamp, and a random number
	return fmt.Sprintf("%s_%d%d%s", strings.TrimSuffix(fileHeader.Filename, ext), time.Now().UnixNano(), rand.Intn(10000), ext)
}
