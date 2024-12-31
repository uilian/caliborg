package tests

import (
	"CalibreMetadataOrganizer/internal/api"
	"testing"
)

func TestFetchMetadata(t *testing.T) {
	description, categories := api.FetchMetadata("Clean Code")
	if description == "" {
		t.Error("Expected a non-empty description")
	}
	if len(categories) == 0 {
		t.Error("Expected at least one category")
	}
}
