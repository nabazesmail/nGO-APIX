// initializers/database.go
package initializers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // Export the DB variable

func ConnectToDB() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to MySQL:", err)
	}
	log.Println("database connected!")
	DB = db // Assign the DB instance to the exported variable
}
