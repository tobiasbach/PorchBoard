package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	url := fmt.Sprintf("http://dataservice.accuweather.com/forecasts/v1/hourly/12hour/1026032?apikey=%s&language=de&details=true&metric=true", os.Getenv("APIKeyAccuWeather"))
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}

// http://dataservice.accuweather.com/forecasts/v1/hourly/12hour/{locationKey}
// location key for 15370 Fredersdorf: '1026032'
// temperature value is accurate according to our heating system temp sensor

// AccuWeather provides hourly forecast up to 12hrs
// AccuWesther provides daily forecast up to 5 days
// Limited to 50 calls a day ~ 1 call every 30 minutes
