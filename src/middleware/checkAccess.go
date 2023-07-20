// middleware/check_access.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabazesmail/gopher/src/models"
)

// CheckAccess is a middleware that checks if the user has the required role to access the route.
func CheckAccess(requiredRole models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user from the context (assuming you have set it in a previous middleware)
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
			c.Abort()
			return
		}

		// Type assertion to get the user as models.User
		u, ok := user.(*models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type in context"})
			c.Abort()
			return
		}

		// Check if the user is an admin or has the required role
		if u.Role == models.Admin || u.Role == requiredRole {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied."})
			c.Abort()
			return
		}
	}
}
