package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hemi519/activity/config"
	"github.com/hemi519/activity/data"
)

func main() {
	cfg, err := config.NewConfig("config.json")
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return
	}

	// Start the periodic data fetching
	go startPeriodicFetching(cfg)

	// Wait for termination signal
	waitForTerminationSignal()
}

func startPeriodicFetching(cfg *config.Config) {
	fetchInterval, err := time.ParseDuration(cfg.FetchInterval)
	if err != nil {
		fmt.Printf("Failed to parse fetch interval: %v\n", err)
		return
	}

	ticker := time.NewTicker(fetchInterval)

	// Fetch and save data on each tick
	for range ticker.C {
		if err := data.FetchAndSaveData(cfg); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func waitForTerminationSignal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
	fmt.Println("Termination signal received. Exiting...")
}
