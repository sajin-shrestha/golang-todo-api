package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sajin-shrestha/golang-todo-api/pkg/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Continuing with system environment variables.")
	}

	db := database.NewDB()
	_ = db
	log.Printf("Connected to database '%s'.\n", os.Getenv("DB_NAME"))

	// simple handler for test
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from backend API!"))
	})

	port := os.Getenv("PORT")
	if port == "" {
    	log.Fatalf("Required environment variable 'PORT' is not set")
	}

	log.Printf("Server listening on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
    	log.Fatalf("Failed to start server: %v", err)
	}
}
