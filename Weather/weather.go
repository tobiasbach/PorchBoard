package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// returns specified weather forecast data queried from AccuWeather API
func main() {
	godotenv.Load("../.env")
	// location key for 15370 Fredersdorf: '1026032'
	url := fmt.Sprintf("http://dataservice.accuweather.com/forecasts/v1/hourly/12hour/1026032?apikey=%s&language=de&details=true&metric=true", os.Getenv("APIKeyAccuWeather"))
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}

// AccuWeather provides hourly forecast up to 12hrs
// https://developer.accuweather.com/accuweather-forecast-api/apis/get/forecasts/v1/hourly/12hour/%7BlocationKey%7D

// AccuWeather provides daily forecast up to 5 days
// https://developer.accuweather.com/accuweather-forecast-api/apis/get/forecasts/v1/daily/5day/%7BlocationKey%7D

// Free plan is limited to 50 calls a day ~ 1 call every 30 minutes
