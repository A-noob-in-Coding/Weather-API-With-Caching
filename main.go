package main

import (
	"WeatherAPI/utils"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage WeatherAPI <location>")
		return
	}
	location := args[1]
	cache := utils.CheckInCache(location)
	if cache != "" {
		utils.PrintResult([]byte(cache))
		return
	}
	content, err := os.ReadFile("api.key")
	if err != nil {
		log.Fatal("Please provide api key in file api.key")
		return
	}
	apiKey := strings.TrimSpace(string(content))

	var url = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + location + "/today?unitGroup=metric&include=current&key=" + apiKey + "&contentType=json"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal("could not fetch weather report")
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	utils.PrintResult(body)
	utils.StoreInCache(body, location)
}
