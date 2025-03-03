package services

import (
	"3rd-party/models"
	"encoding/json"
	"fmt"
	"net/http"
)

const openWeatherAPI = "https://api.openweathermap.org/data/2.5/weather"

// FetchWeather interacts with the OpenWeather API to fetch weather data for a given city
func FetchWeather(city string, apiKey string) (*models.WeatherResponse, error) {
	// Build the API request URL
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", openWeatherAPI, city, apiKey)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather: %s", resp.Status)
	}

	// Parse the response JSON into the WeatherResponse model
	var weatherData models.WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}
