package handlers

import (
	"3rd-party/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY") // Ensure your API key is set as an environment variable
	if apiKey == "" {
		http.Error(w, "API key is not set", http.StatusInternalServerError)
		return
	}

	// Get the 'city' query parameter
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City is required", http.StatusBadRequest)
		return
	}

	// Fetch weather data using the service
	weatherData, err := services.FetchWeather(city, apiKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching weather: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the weather data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}
