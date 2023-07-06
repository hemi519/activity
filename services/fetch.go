package services

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hemi519/activity/config"
)

func StartPeriodicFetching(cfg *config.Config, logger *log.Logger) {
	fetchInterval, err := time.ParseDuration(cfg.FetchInterval)
	if err != nil {
		logger.Printf("Failed to parse fetch interval: %v\n", err)
		return
	}

	ticker := time.NewTicker(fetchInterval)

	// Fetch and save data
	for range ticker.C {
		err := FetchAndSaveData(cfg)
		if err != nil {
			logger.Println("Error:", err)
		} else {
			logger.Println("Data fetched and saved successfully")
		}
	}
}

// this function is typically used in long-running server applications or daemons
// that need to handle graceful shutdowns
func WaitForTerminationSignal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
	fmt.Println("Termination signal received. Exiting...")
}
