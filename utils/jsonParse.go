package utils

import (
	_ "encoding/json"
	"fmt"
	_ "os"

	"github.com/tidwall/gjson"
)

func PrintResult(body []byte) {
	json := string(body)

	current := gjson.Get(json, "currentConditions")

	fmt.Println("======== Current Weather ========")
	fmt.Printf("Location:    %s\n", gjson.Get(json, "resolvedAddress"))
	fmt.Printf("Timezone:    %s\n", gjson.Get(json, "timezone"))
	fmt.Printf("Time:        %s\n", current.Get("datetime"))
	fmt.Printf("Temp:        %.1f°C\n", current.Get("temp").Float())
	fmt.Printf("Feels Like:  %.1f°C\n", current.Get("feelslike").Float())
	fmt.Printf("Humidity:    %.1f%%\n", current.Get("humidity").Float())
	fmt.Printf("Conditions:  %s\n", current.Get("conditions"))
	fmt.Printf("Wind:        %.1f km/h @ %.0f°\n", current.Get("windspeed").Float(), current.Get("winddir").Float())
	fmt.Printf("Pressure:    %.1f hPa\n", current.Get("pressure").Float())
	fmt.Printf("UV Index:    %.1f\n", current.Get("uvindex").Float())
	fmt.Printf("Sunrise:     %s\n", current.Get("sunrise"))
	fmt.Printf("Sunset:      %s\n", current.Get("sunset"))
	fmt.Println("=================================")
}
