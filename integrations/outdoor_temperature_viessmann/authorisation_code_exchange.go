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

// returns a new bearer token using the 'persistent' refresh token
func refreshBearerToken() {
	godotenv.Load(".env")

	url := fmt.Sprintf("https://iam.viessmann.com/idp/v3/token?client_id=%s&grant_type=refresh_token&refresh_token=%s", os.Getenv("ViessmannClientID"), os.Getenv("ViessmannRefreshToken"))

	request, err := http.NewRequest("POST", url, nil)
	// header is necessary for a successful API call
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
	var bearerToken AuthorisationCodeExchangeResponse
	json.Unmarshal([]byte(body), &bearerToken)

	// write bearerToken to .env file
	os.Setenv("ViessmannBearerToken", bearerToken.AccessToken)
}
