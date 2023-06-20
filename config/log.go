package config

import (
	"fmt"
	"log"
	"os"
)

func SetupLogFile() (*os.File, error) {
	logFile, err := os.OpenFile("logs/activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %v", err)
	}

	log.SetOutput(logFile)
	return logFile, nil
}
