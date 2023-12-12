package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("/app/.env"); err != nil {
		log.Fatalf("Faild to load env file: %v", err)
	}
}

func BitflyerApiKey() string {
	return os.Getenv("BITFLYER_API_KEY")
}

func BitflyerApiSecret() string {
	return os.Getenv("BITFLYER_API_SECRET")
}

func LogFileName() string {
	return os.Getenv("LOG_FILE_NAME")
}
