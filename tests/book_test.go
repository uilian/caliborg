package tests

import (
	"CalibreMetadataOrganizer/internal/data"
	"os"
	"testing"
)

func TestLoadBooks(t *testing.T) {
	books := data.LoadBooks("test_library.json") // Ensure this file exists for testing
	if len(books) == 0 {
		t.Error("Expected books to be loaded, got 0")
	}
}

func TestSaveBooks(t *testing.T) {
	books := []data.Book{
		{ID: 1, Title: "Test Book", Tags: []string{"test"}},
	}
	filename := "test_output.json"
	data.SaveBooks(books, filename)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("Expected file to be created")
	}
	os.Remove(filename) // Clean up test file
}
