package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// returns specific weather forecast data queried from AccuWeather API
func Forecast() {
	godotenv.Load(".env")

	// location key for 15370 Fredersdorf: '1026032'
	// TO-DO move to env file

	url := fmt.Sprintf("http://dataservice.accuweather.com/forecasts/v1/hourly/12hour/1026032?apikey=%s&language=en&details=true&metric=true", os.Getenv("APIKeyAccuWeather"))
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	// convert response body to bytes:
	body, _ := ioutil.ReadAll(resp.Body)

	// convert bytes to an array to access single forecast objects
	// 		https://stackoverflow.com/a/53538261
	var val []map[string]interface{}
	if err := json.Unmarshal([]byte(body), &val); err != nil {
		panic(err)
	}

	// iterate over array and print value for key 'IconPhrase'
	// 		fmt.Println(reflect.TypeOf(val)) >> []map[string]interface {}
	// 		map[key-type]val-type
	for i, s := range val {
		if i == 0 {
		} else if i == 1 {
			fmt.Printf("Weather conditions in %s hour: %s\n", fmt.Sprint(i), s["IconPhrase"])
		} else {
			fmt.Printf("Weather conditions in %s hours: %s\n", fmt.Sprint(i), s["IconPhrase"])
		}
	}
}
