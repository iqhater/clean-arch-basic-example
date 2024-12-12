package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// InitEnv wrapper function to get a values from environment
func InitEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error loading %s file\n", filename)
	}
}

// Config struct contains init environments
type Config struct {
	HTTP_PORT   string
	DB_FILENAME string
}

// NewConfig function returns inited server configuration
func NewConfig() *Config {

	HTTP_PORT := getEnv("HTTP_PORT")
	DB_FILENAME := getEnv("DB_FILENAME")

	return &Config{
		HTTP_PORT,
		DB_FILENAME,
	}
}

// getEnv wrapper function to get a value from environment
func getEnv(key string) string {

	value := os.Getenv(key)

	if value == "" {
		log.Printf("Env variable %s does not exist!\n", key)
		return ""
	}
	return value
}
