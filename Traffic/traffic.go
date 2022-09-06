package traffic

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Routes []struct {
	Summary Summary
}

type Summary struct {
	LengthInMeters      int `json: "lengthInMeters"`
	TravelTimeInSeconds int `json: "travelTimeInSeconds"`
}

func Conditions() {
	godotenv.Load(".env")
	versionNumber := "1"
	destination := "52.57596,13.87884" // Amiceria Strausberg
	routePlanningLocations := fmt.Sprintf("%s:%s", os.Getenv("Home"), destination)

	url := fmt.Sprintf("https://api.tomtom.com/routing/%s/calculateRoute/%s/json?key=%s", versionNumber, routePlanningLocations, os.Getenv("APIKeyTomTom"))
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}
