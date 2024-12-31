package categorizer

import (
	"regexp"
	"strings"

	"CalibreMetadataOrganizer/internal/data"
)

func Categorize(title, description string) string {
	title = strings.ToLower(title)
	description = strings.ToLower(description)

	for category, keywords := range data.Categories {
		for _, keyword := range keywords {
			if match, _ := regexp.MatchString("\\b"+keyword+"\\b", title+" "+description); match {
				return category
			}
		}
	}
	return "Uncategorized"
}
