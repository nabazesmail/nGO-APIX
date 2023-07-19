package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nabazesmail/gopher/src/models"
	"github.com/nabazesmail/gopher/src/services"
)

func CreateUser(c *gin.Context) {
	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create the user using the services package
	user, err := services.CreateUser(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"user": user,
	})
}
