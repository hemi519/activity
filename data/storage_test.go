package data_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/hemi519/activity/data"
)

func TestSaveData(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := ioutil.TempFile("", "test_data.json")
	if err != nil {
		t.Fatal("Failed to create temporary file:", err)
	}
	defer os.Remove(tmpfile.Name())

	// Test data
	activities := []*data.Activity{
		{
			Activity: "Rearrange and organize your room",
			Type:     "busywork",
			Price:    0,
		},
		{
			Activity: "Learn how to use a french press",
			Type:     "recreational",
			Price:    0.3,
		},
	}

	// Save the data to the temporary file
	if err := data.SaveData(activities, tmpfile.Name()); err != nil {
		t.Fatal("Failed to save data:", err)
	}

	// Read the content of the temporary file
	content, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal("Failed to read file:", err)
	}

	// Expected output
	expected := `[{"activity":"Rearrange and organize your room","type":"busywork","price":0},{"activity":"Learn how to use a french press","type":"recreational","price":0.3}]`

	// Compare the content of the file with the expected output
	if string(content) != expected {
		t.Errorf("Content does not match.\nExpected:\n%s\nGot:\n%s", expected, string(content))
	}
}
