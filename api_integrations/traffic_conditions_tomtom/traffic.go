package traffic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// returns travel times to specific destinations
func Conditions() {
	godotenv.Load(".env")

	// TO-DO move vars to env file
	versionNumber := "1"
	destination := "52.57596,13.87884" // Strausberg Ferry
	routePlanningLocations := fmt.Sprintf("%s:%s", os.Getenv("Home"), destination)

	url := fmt.Sprintf("https://api.tomtom.com/routing/%s/calculateRoute/%s/json?key=%s", versionNumber, routePlanningLocations, os.Getenv("APIKeyTomTom"))
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	// convert response body to bytes:
	body, _ := ioutil.ReadAll(resp.Body)

	// unmarshal bytes to struct
	var traffic Traffic
	json.Unmarshal([]byte(body), &traffic)

	// return selected values from struct (return = print as long as this is a command line tool)
	fmt.Printf("\n%s minutes to Strausberg Ferry by car\n\n", fmt.Sprint(traffic.Routes[0].Summary.TravelTimeInSeconds/60))
}
