package categorizer

import (
	"regexp"
	"strings"

	"CalibreMetadataOrganizer/internal/data"
)

func Categorize(title, description string) []string {
	title = strings.ToLower(title)
	description = strings.ToLower(description)

	var matches []string
	for category, keywords := range data.Categories {
		for _, keyword := range keywords {
			if match, _ := regexp.MatchString("\\b"+keyword+"\\b", title+" "+description); match {
				matches = append(matches, keyword)
				matches = append(matches, category)
				// Return immediately if a match is found for this category
				return matches
			}
		}
	}
	return matches
}
