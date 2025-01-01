package utils

import (
	"log"
	"os"
)

func Debug(msg string) {
	debug, _ := os.LookupEnv("DEBUG")
	if debug == "true" {
		log.Printf("[DEBUG] %s\n", msg)
	}
}

func Info(msg string) {
	log.Printf("[INFO] %s\n", msg)
}

func Error(msg string) {
	log.Printf("[ERROR] %s\n", msg)
}
