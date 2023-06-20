package services

import (
	"testing"
)

type mockFetcher struct{}

func (m *mockFetcher) FetchData() (*Activity, error) {
	return &Activity{
		Activity: "Test Activity",
		Type:     "Test Type",
		Price:    1.23,
	}, nil
}

func TestFetchData(t *testing.T) {
	fetcher := &mockFetcher{}
	activity, err := fetcher.FetchData()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Assert the fetched activity matches the expected dummy data
	expectedActivity := &Activity{
		Activity: "Test Activity",
		Type:     "Test Type",
		Price:    1.23,
	}

	if !isEqualActivity(activity, expectedActivity) {
		t.Errorf("Expected activity %v, but got %v", expectedActivity, activity)
	}
}

func isEqualActivity(activity1, activity2 *Activity) bool {
	return activity1.Activity == activity2.Activity &&
		activity1.Type == activity2.Type &&
		activity1.Price == activity2.Price
}
