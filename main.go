package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiKey = ""

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	http.HandleFunc("/weather", getWeather)
	http.ListenAndServe(":8080", nil)
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var weatherData WeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temperature := weatherData.Main.Temp
	temperatureCelsius := temperature - 273.15 

	fmt.Fprintf(w, "Current temperature in %s: %.2f°C", city, temperatureCelsius)
}
