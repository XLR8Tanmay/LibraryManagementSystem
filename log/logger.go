package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Init() {
	currentDate := time.Now().Format("2006-01-01")
	logFileName := fmt.Sprintf("libraryManagement-%s.log", currentDate)

	// Open the log file for writing. Create it if it doesn't exist, append to it if it does.
	logFile, err := os.OpenFile("storage/logs/"+logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file %s: %s", logFileName, err)
	}

	// Set the log output to the log file
	log.SetOutput(logFile)
}

func Log(args ...interface{}) {
	log.Println(args...)
}
