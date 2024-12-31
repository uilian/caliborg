package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			Description string   `json:"description"`
			Categories  []string `json:"categories"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

func FetchMetadata(title string) (string, []string) {
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")
	if apiKey == "" {
		fmt.Println("GOOGLE_BOOKS_API_KEY not found in environment")
		return "", nil
	}

	query := fmt.Sprintf(
		"https://www.googleapis.com/books/v1/volumes?q=intitle:%s&key=%s&maxResults=1",
		url.QueryEscape(title), apiKey,
	)

	resp, err := http.Get(query)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("Error fetching data:", err)
		time.Sleep(3 * time.Second)
		return "", nil
	}
	defer resp.Body.Close()

	var apiResponse GoogleBooksResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		fmt.Println("Error decoding API response:", err)
		return "", nil
	}

	if len(apiResponse.Items) > 0 {
		return apiResponse.Items[0].VolumeInfo.Description, apiResponse.Items[0].VolumeInfo.Categories
	}
	return "", nil
}
