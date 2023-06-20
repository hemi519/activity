package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hemi519/activity/config"
)

func SaveData(activities []*Activity, filePath string) error {
	dirPath := filepath.Dir(filePath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
	}

	// Read existing data from the file, if it exists
	existingData, err := ioutil.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	var existingActivities []*Activity
	if len(existingData) > 0 {
		if err := json.Unmarshal(existingData, &existingActivities); err != nil {
			return err
		}
	}

	// Append new activities to the existing activities
	activities = append(existingActivities, activities...)

	// Write the updated activities to the file
	encodedData, err := json.MarshalIndent(activities, "", "\t")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filePath, encodedData, 0644); err != nil {
		return err
	}

	return nil
}

func FetchAndSaveData(cfg *config.Config) error {
	activity, err := FetchData(cfg.BoredAPI)
	if err != nil {
		fmt.Println("Error:", err)
		return fmt.Errorf("error fetching data: %w", err)
	}

	activities := []*Activity{activity}

	if err := SaveData(activities, cfg.Output); err != nil {
		return fmt.Errorf("error saving data: %w", err)
	}

	fmt.Println("Data saved successfully.")

	return nil
}
