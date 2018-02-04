package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var _ = loadConfig()

func loadConfig() error {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return nil
}

var Addr string
var Env string

func init() {
	Addr = os.Getenv("Addr")
	Env = os.Getenv("Env")
	log.Println("Loaded config.")
}
