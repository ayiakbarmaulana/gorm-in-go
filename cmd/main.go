package main

import (
	server "gorm-in-go/internal/api/http"
	"gorm-in-go/pkg/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Initialize database
	_, err := database.InitDB(
		os.Getenv("DB_HOST"),     // host
		os.Getenv("DB_USER"),     // user
		os.Getenv("DB_PASSWORD"), // password
		os.Getenv("DB_NAME"),     // database name
		os.Getenv("DB_PORT"),     // port
	)

	if err != nil {
		log.Fatal(err)
	}

	server.StartServer()
}
