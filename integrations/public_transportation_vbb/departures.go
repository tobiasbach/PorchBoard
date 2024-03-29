package vbb

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// returns next departures from a specific stop within a time span (default 60 minutes from now)
func Departures() string {
	godotenv.Load(".env")

	// calling departureBoard endpoint - requires to get the stopId first
	url := fmt.Sprintf("%s%s%s&%s", os.Getenv("vbbBaseUrl"), os.Getenv("vbbEndpointDepartureBoard"), os.Getenv("vbbStopId"), os.Getenv("vbbAccessId"))

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// convert response body to bytes:
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal bytes to struct
	var departureBoard DepartureBoard
	xml.Unmarshal([]byte(body), &departureBoard)

	// return departures (only the first entry of departureBoard for now)
	value := fmt.Sprintf("Next departure from %s at %s direction %s", fmt.Sprintf(departureBoard.Departure[0].Stop), fmt.Sprint(departureBoard.Departure[0].Time), fmt.Sprintf(departureBoard.Departure[0].Direction))
	return value
}
