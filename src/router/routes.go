// router/router.go
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nabazesmail/gopher/src/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Create a route to create a new user
	r.POST("/users", controllers.CreateUser)

	// Create a route to get all users
	r.GET("/users", controllers.GetAllUsers)

	// Create a route to get a user by ID
	r.GET("/users/:id", controllers.GetUserByID)

	// Create a route to update a user by ID
	r.PUT("/users/:id", controllers.UpdateUserByID)

	// Create a route to delete a user by ID
	r.DELETE("/users/:id", controllers.DeleteUserByID)

	// Add a route for user login
	r.POST("/login", controllers.Login)

	return r
}
