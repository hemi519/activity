package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Activity struct {
	Activity string  `json:"activity"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
}

func FetchData(apiURL string) (*Activity, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var activity Activity
	err = json.NewDecoder(resp.Body).Decode(&activity)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &activity, nil
}
