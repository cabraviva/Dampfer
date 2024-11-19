package utils

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Register Logger
var Log = logrus.New()

func InitLogger() {
	// Set log formatter
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	// Set log level
	Log.SetLevel(logrus.InfoLevel)

	// Get the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		Log.Fatal("Error retrieving home directory: ", err)
	}

	// Define the log file path
	logDir := filepath.Join(homeDir, "Dampfer")
	logFilePath := filepath.Join(logDir, "Dampfer.log")

	// Create the directory if it doesn't exist
	err = os.MkdirAll(logDir, 0755)
	if err != nil {
		Log.Fatal("Error creating log directory: ", err)
	}

	// Create or open the log file (e.g., 'Dampfer.log') in append mode
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		Log.Fatal("Error opening log file: ", err)
	}

	// Use io.MultiWriter to write logs to both the console and the log file
	Log.SetOutput(io.MultiWriter(os.Stdout, logFile))
}
