package api

import (
	"CalibreMetadataOrganizer/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

var googleAPIBaseURL = "https://www.googleapis.com/books/v1/volumes"

func SetGoogleAPIBaseURL(url string) {
	googleAPIBaseURL = url
}

type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			Description string   `json:"description"`
			Categories  []string `json:"categories"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

// FetchMetadata fetches metadata for a book title using the Google Books API.
func FetchMetadata(client *http.Client, title string) (string, []string) {
	if client == nil {
		client = http.DefaultClient
	}

	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")
	if apiKey == "" {
		utils.Error("GOOGLE_BOOKS_API_KEY not found in environment")
		return "", nil
	}

	title = url.QueryEscape(utils.CleanBookName(title))
	query := fmt.Sprintf("%s?q=intitle:%s&key=%s&maxResults=1", googleAPIBaseURL, title, apiKey)
	resp, err := client.Get(query)
	if err != nil {
		utils.Error(fmt.Sprintf("Error making request: %s", err))
		return "", nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		utils.Error(fmt.Sprintf("Error fetching data: %s\n", resp.Status))
		return "", nil
	}

	var apiResponse GoogleBooksResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		utils.Error(fmt.Sprintf("Error decoding API response: %s", err))
		return "", nil
	}

	if len(apiResponse.Items) > 0 {
		return apiResponse.Items[0].VolumeInfo.Description, apiResponse.Items[0].VolumeInfo.Categories
	}
	return "", nil
}
