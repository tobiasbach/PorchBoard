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

// returns travel times to specific destinations
func Conditions() {
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

	// return travel time in minutes (return = print as long as this is a command line tool)
	fmt.Printf("\n%s minutes to Strausberg Ferry by car\n\n", fmt.Sprint(traffic.Routes[0].Summary.TravelTimeInSeconds/60))
}
