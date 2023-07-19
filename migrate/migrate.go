package main

import (
	"github.com/nabazesmail/gopher/initializers"
	"github.com/nabazesmail/gopher/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
