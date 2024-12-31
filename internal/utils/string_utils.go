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

func Unique(slice []string) []string {
	keys := make(map[string]bool)
	var result []string
	for _, s := range slice {
		if _, exists := keys[s]; !exists {
			keys[s] = true
			result = append(result, s)
		}
	}
	return result
}

func CleanBookName(bookName string) string {
	baseName := strings.TrimSuffix(bookName, filepath.Ext(bookName))
	reg := regexp.MustCompile(`[_\-]+| +`)
	cleaned := reg.ReplaceAllString(baseName, " ")
	cleaned = strings.TrimSpace(cleaned)

	casesUtil := cases.Title(language.English)
	return casesUtil.String(strings.ToLower(cleaned))
}
