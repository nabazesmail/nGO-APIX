// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nabazesmail/gopher/controllers"
	"github.com/nabazesmail/gopher/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Create a route to create a new post
	r.POST("/users", controllers.CreateUser)

	r.Run()
}
