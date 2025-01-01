package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"CalibreMetadataOrganizer/internal/api"
)

func TestFetchMetadata(t *testing.T) {
	// Mock HTTP response
	mockResponse := `{
		"items": [{
			"volumeInfo": {
				"description": "A comprehensive guide to writing clean code.",
				"categories": ["Programming", "Software Development"]
			}
		}]
	}`

	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.String(), "intitle:Clean+Code") {
			t.Errorf("Unexpected query: %s", r.URL.String())
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Replace the base URL with the mock server's URL
	api.SetGoogleAPIBaseURL(mockServer.URL)

	description, categories := api.FetchMetadata(nil, "Clean Code")
	expectedDescription := "A comprehensive guide to writing clean code."
	expectedCategories := []string{"Programming", "Software Development"}

	if description != expectedDescription {
		t.Errorf("Expected description to be '%s', got '%s'", expectedDescription, description)
	}
	if len(categories) != len(expectedCategories) || categories[0] != expectedCategories[0] || categories[1] != expectedCategories[1] {
		t.Errorf("Expected categories to be %v, got %v", expectedCategories, categories)
	}
}
