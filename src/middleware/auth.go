package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nabazesmail/gopher/src/models"
	"github.com/nabazesmail/gopher/src/repository"
	"github.com/nabazesmail/gopher/src/utils"
)

// AuthMiddleware is a custom middleware that checks if the request contains a valid JWT token.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}

		// Check if the authorization header contains the "Bearer" prefix
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		// Get the token from the authorization header
		tokenString := authHeaderParts[1]

		// Verify the token using the secret key
		claims, err := utils.VerifyJWTToken(tokenString, []byte(os.Getenv("JWT_SECRET_KEY")))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract user information from the token claims and store it in the context
		userIDFloat, ok := claims["sub"].(float64) // Use float64 instead of uint for type assertion
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Convert the userID from float64 to uint
		userID := uint(userIDFloat)

		// Fetch the user from the database using the userID
		user, err := repository.GetUserByID(strconv.FormatUint(uint64(userID), 10)) // Convert uint to string
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
			c.Abort()
			return
		}

		// Set the user in the context
		c.Set("user", user)

		c.Next()
	}
}

// GetUserFromContext is a helper function to extract the user ID from the context.
func GetUserFromContext(c *gin.Context) *models.User {
	if userID, ok := c.Get("userID"); ok {
		if userIDInt, ok := userID.(uint); ok {
			// Convert the userID from uint to string
			user, err := repository.GetUserByID(fmt.Sprintf("%d", userIDInt))
			if err != nil {
				return nil
			}
			return user
		}
	}
	return nil
}
