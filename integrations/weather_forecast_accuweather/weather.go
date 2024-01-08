package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// returns specific weather forecast data queried from AccuWeather API
func Forecast() {
	godotenv.Load(".env")

	url := fmt.Sprintf("http://dataservice.accuweather.com/forecasts/v1/hourly/12hour/%s?apikey=%s&language=en&details=true&metric=true", os.Getenv("AccuWeatherLocationKey"), os.Getenv("AccuWeatherAPIKey"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// convert response body to bytes:
	body, _ := io.ReadAll(resp.Body)

	// unmarshal bytes to array of structs
	var weather []struct{ Weather }
	if err := json.Unmarshal([]byte(body), &weather); err != nil {
		panic(err)
	}

	// iterate over array of structs and return selected values (return = print as long as this is a command line tool)
	// TO-DO: assign strings to variable and return the variable
	for i, w := range weather {
		if i == 0 {
		} else if i == 1 {
			fmt.Printf("Weather conditions in %s hour: %s, %s degrees\n", fmt.Sprint(i), w.IconPhrase, fmt.Sprint(w.Temperature.Value))
		} else {
			fmt.Printf("Weather conditions in %s hours: %s, %s degrees\n", fmt.Sprint(i), w.IconPhrase, fmt.Sprint(w.Temperature.Value))
		}
	}
}
