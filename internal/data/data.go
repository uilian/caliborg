package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Comments string   `json:"comments"`
	Tags     []string `json:"tags"`
}

var Categories = map[string][]string{
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

func LoadBooks(filename string) []Book {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var books []Book
	if err := json.Unmarshal(file, &books); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}
	return books
}

func SaveBooks(books []Book, filename string) {
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		os.Exit(1)
	}
	os.WriteFile(filename, data, 0644)
}
