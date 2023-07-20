// main.go
package main

import (
	"log"
	"os"

	"github.com/nabazesmail/gopher/src/initializers"
	"github.com/nabazesmail/gopher/src/router"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	// Set up custom log format with timestamps
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Optionally redirect logs to a file
	file, err := os.Create("app.log")
	if err != nil {
		log.Fatal("Failed to create log file: ", err)
	}
	log.SetOutput(file)
}

func main() {
	r := router.SetupRouter()
	r.Run()
}
