package utils

import (
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ContainsIgnoreCase(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}

func Unique(slice []string) []string {
	keys := make(map[string]bool)
	var result []string

	for _, s := range slice {
		// Convert to lowercase for comparison
		lower := strings.ToLower(s)
		// Use lowercase strings as keys to check uniqueness
		if _, exists := keys[lower]; !exists {
			keys[lower] = true
			// Add original string (not lowercase) to results
			result = append(result, s)
		}
	}

	return result
}

// CleanBookName cleans the book name by removing underscores, hyphens,
// multiple spaces and removing the file extension from the book name.
// The final returned value is the lowercase title case.
func CleanBookName(bookName string) string {
	baseName := strings.TrimSuffix(bookName, filepath.Ext(bookName))
	reg := regexp.MustCompile(`[_\-]+| +`)
	cleaned := strings.TrimSpace(reg.ReplaceAllString(baseName, " "))
	casesUtil := cases.Title(language.English)
	return casesUtil.String(strings.ToLower(cleaned))
}
