package traffic

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func traffic() {
	godotenv.Load(".env")
	versionNumber := "1"
	destination := "52.575965519202995, 13.878845389884987"
	// Amiceria Strausberg
	routePlanningLocations := fmt.Sprintf("%s:%s", os.Getenv("Home"), destination)
	url := fmt.Sprintf("https://api.tomtom.com/routing/{%s}/calculateRoute/{%s}/{contentType}?key={%s}", versionNumber, routePlanningLocations, os.Getenv("APIKeyAccuWeather"))
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}

// TomTom API
// https://developer.tomtom.com/routing-api/documentation/routing/calculate-route
