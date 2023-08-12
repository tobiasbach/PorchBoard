package temperature

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// returns current outside temperature provided by personal heating system
func Outside() {
	refreshBearerToken()

	godotenv.Load(".env")

	url := fmt.Sprintf("https://api.viessmann.com/iot/v2/features/installations/%s/gateways/%s/devices/%s/features/heating.sensors.temperature.outside", os.Getenv("ViessmannInstallationID"), os.Getenv("ViessmannGatewaySerial"), os.Getenv("ViessmannDeviceID"))

	var bearerToken = "Bearer " + os.Getenv("ViessmannBearerToken")

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", bearerToken)

	client := &http.Client{}
	resp, err := client.Do(request)
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
	var outsideTemperature OutsideTemperature
	json.Unmarshal([]byte(body), &outsideTemperature)

	// return outsideTemperature (return = print as long as this is a command line tool)
	fmt.Printf("\nIt's %.1f degrees outside\n\n", outsideTemperature.Data.Properties.Value.Value)
}
