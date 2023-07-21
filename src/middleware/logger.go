package middleware

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	// Open the app.log file for logging
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open app.log file: %v", err)
	}

	// Create a new logger instance that writes to the logFile
	Logger = log.New(logFile, "", log.Ldate|log.Ltime)
}
