package config

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func createTempConfigFile(data string) (string, error) {
	file, err := ioutil.TempFile("", "config_*.json")
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func removeTempConfigFile(path string) error {
	return os.Remove(path)
}

func TestNewConfig(t *testing.T) {
	// Define the expected values
	expectedFetchInterval := 30 * time.Minute
	expectedOutput := "output.json"

	// Define the test configuration data as a string
	testConfigData := `{
		"fetch_interval": "30m",
		"output": "output.json"
	}`

	// Create a temporary config file with test data
	tempConfigFile, err := createTempConfigFile(testConfigData)
	if err != nil {
		t.Fatalf("Failed to create temporary config file: %v", err)
	}
	defer func() {
		err := removeTempConfigFile(tempConfigFile)
		if err != nil {
			t.Fatalf("Failed to remove temporary config file: %v", err)
		}
	}()

	// Load the config
	cfg, err := NewConfig(tempConfigFile)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Check the fetch interval
	if cfg.FetchInterval != expectedFetchInterval.String() {
		t.Errorf("FetchInterval mismatch. Expected: %v, Got: %v", expectedFetchInterval, cfg.FetchInterval)
	}

	// Check the output file path
	if cfg.Output != expectedOutput {
		t.Errorf("Output mismatch. Expected: %s, Got: %s", expectedOutput, cfg.Output)
	}
}
