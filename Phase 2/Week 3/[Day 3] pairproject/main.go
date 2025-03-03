package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "3rd-party/internal/handler"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// Set up the HTTP server
	http.HandleFunc("/weather", handlers.WeatherHandler)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
