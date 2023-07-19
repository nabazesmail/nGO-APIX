// main.go
package main

import (
	"github.com/nabazesmail/gopher/initializers"
	"github.com/nabazesmail/gopher/router"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := router.SetupRouter()
	r.Run()
}
