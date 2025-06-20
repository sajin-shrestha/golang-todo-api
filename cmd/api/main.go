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
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
				<title>Welcome Page</title>
				<style>
					/* Reset some default styles */
					* {
						margin: 0;
						padding: 0;
						box-sizing: border-box;
					}
					body, html {
						height: 100%;
						font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
						background: linear-gradient(135deg, #6a11cb 0%, #2575fc 100%);
						color: #fff;
						display: flex;
						align-items: center;
						justify-content: center;
					}
					.container {
						background: rgba(255, 255, 255, 0.1);
						padding: 40px 60px;
						border-radius: 12px;
						box-shadow: 0 8px 24px rgba(0, 0, 0, 0.2);
						text-align: center;
						max-width: 400px;
						width: 90%;
					}
					h1 {
						font-size: 2.5rem;
						margin-bottom: 20px;
						letter-spacing: 1.2px;
					}
					p {
						font-size: 1.2rem;
						line-height: 1.6;
						margin-bottom: 30px;
					}
					button {
						background-color: #fff;
						color: #2575fc;
						border: none;
						padding: 12px 30px;
						font-size: 1rem;
						font-weight: 600;
						border-radius: 30px;
						cursor: pointer;
						transition: background-color 0.3s ease, color 0.3s ease;
						box-shadow: 0 4px 15px rgba(255, 255, 255, 0.3);
					}
					button:hover {
						background-color: #1a57d0;
						color: #fff;
					}
					footer {
						margin-top: 40px;
						font-size: 0.9rem;
						color: #d0d0d0;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Welcome to Your Dashboard</h1>
					<p>Hello, Sajin Shrestha! We're glad to see you here.</p>
					<button onclick="alert('Have a great day!')">Click Me</button>
					<footer>© 2025 Sajin's Backend</footer>
				</div>
			</body>
			</html>
		`))
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
