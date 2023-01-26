package vbb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// returns next departures from a specific stop within a time span (default 60 minutes from now)
func Departures() {
	godotenv.Load(".env")

	// calling departureBoard endpoint - requires to get the stopId first
	url := fmt.Sprintf("%s%s%s&%s", os.Getenv("vbbBaseUrl"), os.Getenv("vbbEndpointDepartureBoard"), os.Getenv("vbbStopId"), os.Getenv("vbbAccessId"))
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	// convert response body to bytes:
	body, _ := ioutil.ReadAll(resp.Body)

	// TO-DO fix 400 response - url is double checked and works if pasted into a browser
	// TO-DO XML parsing
	fmt.Print(string(body))
}
