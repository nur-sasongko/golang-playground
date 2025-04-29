package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const baseURL = "https://api.openweathermap.org/data/2.5/weather"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go [city-name]")
	}

	city := os.Args[1]
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	if apiKey == "" {
		log.Fatal("Please set your OPENWEATHER_API_KEY environment variable.")
	}

	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", baseURL, city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Request failed:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API error: %s\n", resp.Status)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		log.Fatal("Failed to parse JSON:", err)
	}

	fmt.Printf("📍 %s\n🌡️  Temp: %.1f°C\n🌤️  Condition: %s\n",
		weather.Name, weather.Main.Temp, weather.Weather[0].Description)
}
