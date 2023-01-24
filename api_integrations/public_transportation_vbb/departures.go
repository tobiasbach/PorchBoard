package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/joho/godotenv"
)

// returns next departures from a specific stop within a time span (default 60 minutes from now)
func departures() {
	godotenv.Load(".env")

	url := fmt.Sprintf("")
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	// convert response body to bytes:
	body, _ := ioutil.ReadAll(resp.Body)
}
