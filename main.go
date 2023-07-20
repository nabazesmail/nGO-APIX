// main.go
package main

import (
	"github.com/nabazesmail/gopher/src/migrate"
	"github.com/nabazesmail/gopher/src/router"
)

// Run the migration logic
func init() {
	migrate.Migration()
}

func main() {
	r := router.SetupRouter()
	r.Run()
}
