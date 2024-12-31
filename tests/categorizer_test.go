package tests

import (
	"CalibreMetadataOrganizer/internal/categorizer"
	"testing"
)

func TestCategorize(t *testing.T) {
	title := "Learning Python"
	description := "An introduction to programming with Python."
	category := categorizer.Categorize(title, description)

	if category != "Programming" {
		t.Errorf("Expected category 'Programming', got '%s'", category)
	}

	title = "A historical account of the medieval era"
	description = "A book about ancient civilizations and medieval history."
	category = categorizer.Categorize(title, description)

	if category != "History" {
		t.Errorf("Expected category 'History', got '%s'", category)
	}
}
