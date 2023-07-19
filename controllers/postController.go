package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nabazesmail/gopher/initializers"
	"github.com/nabazesmail/gopher/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var body struct {
		FullName string        `json:"full_name" binding:"required"`
		Username string        `json:"username" binding:"required,alphanum"`
		Password string        `json:"password" binding:"required,min=8,max=15"`
		Status   models.Status `json:"status" binding:"omitempty,oneof=active inactive"`
		Role     models.Role   `json:"role" binding:"omitempty,oneof=admin operator"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// If status or role not provided in the request, set default values
	if body.Status == "" {
		body.Status = models.Active // Set default status to "active" if not provided
	}
	if body.Role == "" {
		body.Role = models.Operator // Set default role to "operator" if not provided
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Status(500)
		return
	}

	// Create a new User instance with the hashed password
	user := models.User{
		FullName: body.FullName,
		Username: body.Username,
		Password: string(hashedPassword),
		Status:   body.Status,
		Role:     body.Role,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
