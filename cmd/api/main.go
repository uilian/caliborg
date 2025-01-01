package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"CalibreMetadataOrganizer/internal/api"
	"CalibreMetadataOrganizer/internal/categorizer"
	"CalibreMetadataOrganizer/internal/data"
	"CalibreMetadataOrganizer/internal/utils"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		utils.Error("Error loading .env file")
		os.Exit(1)
	}
	// Load from the environment variables the path file
	// containing the saved library contents.
	libraryPath, ok := os.LookupEnv("LIBRARY_PATH")
	if !ok {
		utils.Error("LIBRARY_PATH not found in environment")
		os.Exit(1)
	}

	utils.Debug("Loading library from " + libraryPath)

	books := data.LoadBooks(libraryPath)

	var updatedBooks []data.Book
	var commands []string
	const rateLimit = 100 // set to the max number of calls per minute allowed by the API.
	minInterval := time.Minute / rateLimit

	for _, book := range books {
		// Sleep for the minimum interval to avoid API rate limiting.
		time.Sleep(minInterval)

		description, apiCategories := api.FetchMetadata(nil, utils.CleanBookName(book.Title))
		categories := categorizer.Categorize(book.Title, description)

		if len(apiCategories) > 0 {
			book.Tags = append(book.Tags, apiCategories...)
		}
		if len(categories) > 0 {
			book.Tags = append(book.Tags, categories...)
		}
		// Remove duplicates
		book.Tags = utils.Unique(book.Tags)

		updatedBooks = append(updatedBooks, book)
		data.SaveBooks(updatedBooks, "partial_results.json") // Save progress

		command := fmt.Sprintf(`calibredb set_metadata --ids %d --tags "%s"`, book.ID, strings.Join(book.Tags, ", "))
		commands = append(commands, command)
	}

	data.SaveBooks(updatedBooks, "updated_library_with_api.json")
	os.WriteFile("update_tags_with_api.sh", []byte(strings.Join(commands, "\n")), 0755)
	utils.Info("Categorization with API complete. Generated update_tags_with_api.sh.")
}
