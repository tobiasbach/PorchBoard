package car

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// all functions in this package activate a 'shelly scene' via webhook
// these scenes have been created via shelly app / web ui
// the current setup features a shelly plus 2 pm connected to
// output 1 = engine preheating
// output 2 = cabin heating fan
// each scene activates one or both outputs for 30 minutes or disables all of them

func WarmUpEngineAndCabin() string {
	godotenv.Load(".env")

	url := "https://shelly-89-eu.shelly.cloud/scene/manual_run&id=SCENE_ID&auth_key=AUTH_KEY"

	request, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	value := "Engine and cabin are being warmed up."
	return value
}

func WarmUpEngineOnly() string {
	value := "Engine is being warmed up."
	return value
}

func WarmUpCabinOnly() string {
	value := "Cabin is being warmed up."
	return value
}

func QuitWarmUp() string {
	value := "Warm up has been stopped."
	return value
}
