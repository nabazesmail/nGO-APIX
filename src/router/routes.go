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

	return r
}
