package main

import (
	"github.com/nabazesmail/gopher/src/initializers"
	"github.com/nabazesmail/gopher/src/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
