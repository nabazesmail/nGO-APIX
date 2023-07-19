// main.go
package main

import (
	"github.com/nabazesmail/gopher/src/initializers"
	"github.com/nabazesmail/gopher/src/router"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := router.SetupRouter()
	r.Run()
}
