package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hemi519/activity/config"
)

func SaveData(activities []*Activity, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encodedData, err := json.MarshalIndent(activities, "", "\t")
	if err != nil {
		return err
	}

	_, err = file.Write(encodedData)
	if err != nil {
		return err
	}

	return nil
}

func FetchAndSaveData(cfg *config.Config) error {
	activity, err := FetchData()
	if err != nil {
		fmt.Println("Error:", err)
		return fmt.Errorf("error fetching data: %w", err)
	}

	activities := []*Activity{activity}

	// Load existing data from the file, if it exists
	if _, err = os.Stat(cfg.Output); err == nil {
		existingData, err := ioutil.ReadFile(cfg.Output)
		if err != nil {
			return fmt.Errorf("error read data: %w", err)
		}

		var existingActivities []*Activity
		err = json.Unmarshal(existingData, &existingActivities)
		if err != nil {

			return fmt.Errorf("json Unmarshal Error: %w", err)
		}

		activities = append(activities, existingActivities...)
	}

	if err := SaveData(activities, cfg.Output); err != nil {
		return fmt.Errorf("error saving data: %w", err)
	}

	fmt.Println("Data saved successfully.")

	return nil
}
