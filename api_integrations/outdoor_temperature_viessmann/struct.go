package temperature

type OutsideTemperature struct {
	Data struct {
		APIVersion int64    `json:"apiVersion"`
		Commands   struct{} `json:"commands"`
		DeviceID   string   `json:"deviceId"`
		Feature    string   `json:"feature"`
		GatewayID  string   `json:"gatewayId"`
		IsEnabled  bool     `json:"isEnabled"`
		IsReady    bool     `json:"isReady"`
		Properties struct {
			Status struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"status"`
			Value struct {
				Type  string  `json:"type"`
				Unit  string  `json:"unit"`
				Value float64 `json:"value"`
			} `json:"value"`
		} `json:"properties"`
		Timestamp string `json:"timestamp"`
		URI       string `json:"uri"`
	} `json:"data"`
}
