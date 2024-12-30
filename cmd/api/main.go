package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Book struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Comments string   `json:"comments"`
	Tags     []string `json:"tags"`
}

type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			Description string   `json:"description"`
			Categories  []string `json:"categories"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

var categories = map[string][]string{
	"Programming": {
		"python", "java", "scala", "javascript", "coding", "software", "golang", "typescript",
		"rust", "kotlin", "ruby", "c++", "c#", "php", "swift", "html", "css", "sql", "bash",
		"docker", "kubernetes", "react", "angular", "vue", "django", "flask", "spring", "express",
		"nodejs", "tensorflow", "pytorch", "machine learning", "artificial intelligence",
		"data science", "big data", "hadoop", "spark", "haskell", "elixir", "assembly",
		"blockchain", "web development", "backend", "frontend", "full stack", "devops",
		"cloud computing", "aws", "azure", "gcp",
	},
	"Fiction": {
		"novel", "fiction", "literature", "story", "drama", "adventure", "classic",
		"short story", "narrative", "fictional", "prose",
	},
	"History": {
		"history", "biography", "historical", "ancient", "medieval", "renaissance", "modern",
		"world war", "civil war", "revolution", "historical fiction", "prehistoric", "dynasty",
		"empire", "colonial", "industrial", "enlightenment", "archaeology",
	},
	"Science": {
		"physics", "chemistry", "biology", "science", "technology", "astronomy", "genetics",
		"quantum", "nanotechnology", "robotics", "space exploration", "ecology", "geology",
		"meteorology", "zoology", "botany", "mathematics", "statistics", "engineering",
	},
	"Fantasy": {
		"fantasy", "magic", "dragons", "sorcery", "epic", "myth", "legend", "fairy tale",
		"elves", "wizards", "dwarves", "orcs", "enchantment", "adventure", "alternate world",
		"quest", "dark fantasy", "high fantasy", "urban fantasy",
	},
	"Romance": {
		"romance", "love", "relationship", "affair", "passion", "heartbreak", "romantic",
		"couple", "dating", "soulmate", "chemistry", "wedding", "marriage", "honeymoon",
		"romantic comedy", "historical romance",
	},
	"Mystery": {
		"mystery", "detective", "thriller", "crime", "investigation", "suspense", "whodunit",
		"murder", "forensics", "spy", "heist", "conspiracy", "puzzle", "private eye",
		"clues", "cold case", "noir", "psychological thriller",
	},
	"Horror": {
		"horror", "ghost", "vampire", "werewolf", "haunted", "supernatural", "zombie",
		"paranormal", "occult", "monster", "slasher", "gothic", "dark", "terror",
		"psychological horror",
	},
	"Philosophy": {
		"philosophy", "ethics", "metaphysics", "existentialism", "logic", "epistemology",
		"stoicism", "nihilism", "aesthetics", "rationalism", "idealism", "realism",
		"philosophical", "ontology",
	},
	"Health": {
		"health", "fitness", "nutrition", "medicine", "mental health", "wellness",
		"exercise", "yoga", "diet", "therapy", "self-care", "healthcare", "anatomy",
		"psychology", "public health",
	},
	"Uncategorized": {},
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	file, err := os.ReadFile("library.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var books []Book
	err = json.Unmarshal(file, &books)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	var updatedBooks []Book
	var commands []string

	// Add rate limiting variables
	const rateLimit = 100 // calls per minute
	minInterval := time.Minute / rateLimit

	for _, book := range books {
		// Add rate limiting delay
		time.Sleep(minInterval)
		description, apiCategories := fetchMetadata(book.Title)
		category := categorizeBook(book.Title, description)
		if !contains(book.Tags, category) {
			book.Tags = append(book.Tags, category)
		}
		book.Tags = append(book.Tags, apiCategories...)
		book.Tags = unique(book.Tags)

		updatedBooks = append(updatedBooks, book)
		command := fmt.Sprintf(`calibredb set_metadata --ids %d --tags "%s"`, book.ID, strings.Join(book.Tags, ", "))
		commands = append(commands, command)
	}

	updatedFile, err := json.MarshalIndent(updatedBooks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling updated books:", err)
		os.Exit(1)
	}
	os.WriteFile("updated_library_with_api.json", updatedFile, 0644)

	os.WriteFile("update_tags_with_api.sh", []byte(strings.Join(commands, "\n")), 0755)
	fmt.Println("Categorization with API complete. Generated update_tags_with_api.sh.")
}

func fetchMetadata(title string) (string, []string) {
	apiKey := os.Getenv("GOOGLE_BOOKS_API_KEY")
	if apiKey == "" {
		fmt.Println("GOOGLE_BOOKS_API_KEY not found in environment")
		return "", nil
	}
	title = url.QueryEscape(cleanBookName(title))
	googleBooksUrl := "https://www.googleapis.com/books/v1/volumes?q=intitle:%s&key=%s&maxResults=1"
	query := fmt.Sprintf(googleBooksUrl, title, apiKey)
	fmt.Println("Fetching metadata for: ", title)
	resp, err := http.Get(query)
	fmt.Println("Response: ", resp.StatusCode)
	// rate limit response handling
	if err != nil || resp.StatusCode == 429 {
		fmt.Println("Rate limit exceeded, waiting before retry...")
		time.Sleep(time.Second * 3) // Wait 3 seconds before potential retry
		return "", nil
	} else if resp.StatusCode != 200 {
		fmt.Println("Error fetching data:", err)
		return "", nil
	}
	defer resp.Body.Close()

	var apiResponse GoogleBooksResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		fmt.Println("Error decoding API response:", err)
		return "", nil
	}

	if len(apiResponse.Items) > 0 {
		return apiResponse.Items[0].VolumeInfo.Description, apiResponse.Items[0].VolumeInfo.Categories
	}
	return "", nil
}

func categorizeBook(title, description string) string {
	title = strings.ToLower(title)
	description = strings.ToLower(description)

	for category, keywords := range categories {
		for _, keyword := range keywords {
			if match, _ := regexp.MatchString("\\b"+keyword+"\\b", title+" "+description); match {
				return category
			}
		}
	}
	return "Uncategorized"
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func unique(slice []string) []string {
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

// cleanBookName cleans up a book name string by removing file extensions,
// replacing undesired characters with spaces, and formatting the result.
func cleanBookName(bookName string) string {
	// Remove the file extension
	baseName := strings.TrimSuffix(bookName, filepath.Ext(bookName))

	// Replace underscores, hyphens, and multiple spaces with a single space
	reg := regexp.MustCompile(`[_\-]+| +`)
	cleaned := reg.ReplaceAllString(baseName, " ")

	// Trim leading and trailing spaces
	cleaned = strings.TrimSpace(cleaned)

	// Capitalize the first letter of each word
	casesUtil := cases.Title(language.English)
	cleaned = casesUtil.String(strings.ToLower(cleaned))

	return cleaned
}
