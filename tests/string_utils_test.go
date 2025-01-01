package tests

import (
	"CalibreMetadataOrganizer/internal/utils"
	"testing"
)

func TestContains(t *testing.T) {
	slice := []string{"a", "b", "c"}
	if !utils.Contains(slice, "a") || !utils.ContainsIgnoreCase(slice, "A") {
		t.Error("Expected 'a' to be in slice")
	}
	if utils.Contains(slice, "d") || utils.ContainsIgnoreCase(slice, "D") {
		t.Error("Expected 'd' to not be in slice")
	}
}

func TestUnique(t *testing.T) {
	slice := []string{"a", "b", "a", "c", "b"}
	expected := []string{"a", "b", "c"}
	result := utils.Unique(slice)
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}
}

func TestCleanBookName(t *testing.T) {
	bookName := "clean-code.pdf"
	expected := "Clean Code"
	result := utils.CleanBookName(bookName)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
