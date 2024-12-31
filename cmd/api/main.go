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
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	books := data.LoadBooks("library.json")

	var updatedBooks []data.Book
	var commands []string
	const rateLimit = 100 // calls per minute
	minInterval := time.Minute / rateLimit

	for _, book := range books {
		time.Sleep(minInterval)
		description, apiCategories := api.FetchMetadata(utils.CleanBookName(book.Title))
		category := categorizer.Categorize(book.Title, description)

		if !utils.Contains(book.Tags, category) {
			book.Tags = append(book.Tags, category)
		}
		book.Tags = append(book.Tags, apiCategories...)
		book.Tags = utils.Unique(book.Tags)

		updatedBooks = append(updatedBooks, book)
		data.SaveBooks(updatedBooks, "partial_results.json") // Save progress

		command := fmt.Sprintf(`calibredb set_metadata --ids %d --tags "%s"`, book.ID, strings.Join(book.Tags, ", "))
		commands = append(commands, command)
	}

	data.SaveBooks(updatedBooks, "updated_library_with_api.json")
	os.WriteFile("update_tags_with_api.sh", []byte(strings.Join(commands, "\n")), 0755)
	fmt.Println("Categorization with API complete. Generated update_tags_with_api.sh.")
}
