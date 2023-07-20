// controllers/controllers.go
package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nabazesmail/gopher/src/models"
	"github.com/nabazesmail/gopher/src/services"
)

func CreateUser(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error parsing request body: %s", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Create the user using the services package
	user, err := services.CreateUser(&body)
	if err != nil {
		log.Printf("Error creating user: %s", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(201, gin.H{
		"user": user,
	})
}

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"users": users})
}

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

func DeleteUserByID(c *gin.Context) {
	userID := c.Param("id")

	err := services.DeleteUserByID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func Login(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error parsing request body: %s", err)
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
