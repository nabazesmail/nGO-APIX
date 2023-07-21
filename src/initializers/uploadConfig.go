package initializers

import (
	"mime/multipart"
	"path/filepath"
	"strings"
)

// Helper function to check if the uploaded file is an image
func IsImageFile(fileHeader *multipart.FileHeader) bool {
	// Extract the file extension from the uploaded file's header
	ext := filepath.Ext(fileHeader.Filename)
	ext = strings.ToLower(ext)
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}
