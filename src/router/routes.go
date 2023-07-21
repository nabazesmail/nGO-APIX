// router/router.go
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nabazesmail/gopher/src/controllers"
	"github.com/nabazesmail/gopher/src/middleware"
	"github.com/nabazesmail/gopher/src/models"
)

// SetupRouter sets up the Gin router and defines the routes for the application.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//  a route to create a new user
	r.POST("/register", controllers.CreateUser)

	//  a route to login the user
	r.POST("/login", controllers.Login)

	//  protected routes using a middleware to authenticate the requests.
	protectedRoutes := r.Group("/")
	protectedRoutes.Use(middleware.AuthMiddleware()) // Use the AuthMiddleware for all routes in this group.

	//  a route to get all users (protected route)
	protectedRoutes.GET("/users", middleware.CheckAccess(models.Operator), controllers.GetAllUsers)

	//  a route to get a user by ID (protected route)
	protectedRoutes.GET("/users/:id", middleware.CheckAccess(models.Operator), controllers.GetUserByID)

	//  a route to update a user by ID (protected route)
	protectedRoutes.PUT("/users/:id", middleware.CheckAccess(models.Admin), controllers.UpdateUserByID)

	//  a route to delete a user by ID (protected route)
	protectedRoutes.DELETE("/users/:id", middleware.CheckAccess(models.Admin), controllers.DeleteUserByID)

	// Create a route to get the user's profile (protected route)
	protectedRoutes.GET("/profile", middleware.CheckAccess(models.Operator), controllers.GetUserProfile)

	// Create a route to handle file uploads and update user profile picture
	protectedRoutes.POST("/imgUpload/:id", middleware.CheckAccess(models.Admin), controllers.UploadProfilePicture)
	return r
}
