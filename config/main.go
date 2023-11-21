package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MONGO_URL      string
	JWT_SECRET     string
	RF_JWT_SECRET  string
	IDENTIFY_KEY   string
	SMTP_USERNAME  string
	SMTP_PASSWORD  string
	BING_MAP_API   string
	GOOGLE_MAP_API string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		panic(err)
	}

	MONGO_URL = os.Getenv("MONGO_URL")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	IDENTIFY_KEY = os.Getenv("IDENTIFY_KEY")
	RF_JWT_SECRET = os.Getenv("RF_JWT_SECRET")
	SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
	SMTP_USERNAME = os.Getenv("SMTP_USERNAME")
	BING_MAP_API = os.Getenv("BING_MAP_API")
	GOOGLE_MAP_API = os.Getenv("GOOGLE_MAP_API")
}
