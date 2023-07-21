// controllers/controllers.go
package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabazesmail/gopher/src/middleware"
	"github.com/nabazesmail/gopher/src/models"
	"github.com/nabazesmail/gopher/src/services"
	"github.com/nabazesmail/gopher/src/utils"
)

// create user
func CreateUser(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		middleware.Logger.Printf("Error parsing request body: %s", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Create the user using the services package
	user, err := services.CreateUser(&body)
	if err != nil {
		middleware.Logger.Printf("Error creating user: %s", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(201, gin.H{
		"user": user,
	})
}

// user login
func Login(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		middleware.Logger.Printf("Error parsing request body: %s", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Check if the username and password are provided
	if body.Username == "" || body.Password == "" {
		c.JSON(400, gin.H{"error": "Username and password must be provided"})
		return
	}

	// Authenticate user using the services package
	token, err := services.AuthenticateUser(&body)
	if err != nil {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

// getting all users
func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"users": users})
}

// getting one user by Id
func GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	user, err := services.GetUserByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if user == nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

// updating user
func UpdateUserByID(c *gin.Context) {
	userID := c.Param("id")

	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := services.UpdateUserByID(userID, &body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if user == nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

// deleting user
func DeleteUserByID(c *gin.Context) {
	userID := c.Param("id")

	err := services.DeleteUserByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

// getting user profile only with token
func GetUserProfile(c *gin.Context) {
	// Extract the user from the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	// Type assertion to get the user as models.User
	u, ok := user.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type in context"})
		return
	}

	// Create a UserResponse struct without the password field
	userResponse := utils.UserResponse{
		ID:       u.ID,
		FullName: u.FullName,
		Username: u.Username,
		Status:   string(u.Status),
		Role:     string(u.Role),
	}

	// Return the user's profile
	c.JSON(http.StatusOK, gin.H{
		"user": userResponse,
	})
}

// uploading profile pic
func UploadProfilePicture(c *gin.Context) {
	userID := c.Param("id")

	// Check if the request contains a file with the key "profile_picture"
	file, fileHeader, err := c.Request.FormFile("profile_picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing file in the request"})
		return
	}
	defer file.Close()

	// Update the user's profile picture
	user, err := services.UpdateUserProfilePicture(userID, fileHeader)
	if err != nil {
		log.Printf("Error updating user's profile picture: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile picture"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// fetching profile pic
func GetProfilePicture(c *gin.Context) {
	userID := c.Param("id")

	// Retrieve the user's profile picture data using the services package
	data, err := services.GetProfilePictureByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch profile picture"})
		return
	}

	if data == nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Determine the content type based on the file extension
	contentType := http.DetectContentType(data)

	// Set the appropriate Content-Type header for image preview
	c.Header("Content-Type", contentType)

	// Copy the profile picture data to the response body for previewing the profile picture
	_, err = c.Writer.Write(data)
	if err != nil {
		middleware.Logger.Printf("Error copying profile picture data: %s", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve profile picture"})
		return
	}
}
