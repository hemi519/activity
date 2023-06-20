package main

import (
	"fmt"
	"log"

	"github.com/hemi519/activity/config"
	"github.com/hemi519/activity/services"
)

func main() {
	//load configs
	cfg, err := config.NewConfig("config.json")
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return
	}

	//log data
	logFile, err := config.SetupLogFile()
	if err != nil {
		fmt.Println("Failed to create log file:", err)
		return
	}
	defer logFile.Close()

	fmt.Println("Application Running.....")

	// Start the periodic data fetching
	go services.StartPeriodicFetching(cfg, log.New(logFile, "", log.LstdFlags))

	// Wait for termination signal
	services.WaitForTerminationSignal()
}
