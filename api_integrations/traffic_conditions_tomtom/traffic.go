package traffic

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// returns travel time to a specific destination
// to-do: to check the actual condition of your preferred path, adjust route by adding more locations
//
//	the current implementation routes from a to b which may returns your common expected travel
//	time but actually uses a different path due to traffic jams, etc
func Conditions() string {
	godotenv.Load(".env")

	route := fmt.Sprintf("%s:%s", os.Getenv("TomTomHomeCoordinates"), os.Getenv("TomTomDestinationCoordinates"))

	url := fmt.Sprintf("https://api.tomtom.com/routing/%s/calculateRoute/%s/json?key=%s", os.Getenv("TomTomAPIVersionNumber"), route, os.Getenv("TomTomAPIKey"))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// convert response body to bytes:
	body, _ := io.ReadAll(resp.Body)

	// unmarshal bytes to struct
	var traffic Traffic
	json.Unmarshal([]byte(body), &traffic)

	// return trafficConditions (travel time in minutes)
	value := fmt.Sprintf("%s minutes to Strausberg ferry by car", fmt.Sprint(traffic.Routes[0].Summary.TravelTimeInSeconds/60))
	return value
}
