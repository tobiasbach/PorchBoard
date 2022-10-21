package main

import (
	traffic "PorchBoard/api_integrations/traffic_conditions_tomtom"
	weather "PorchBoard/api_integrations/weather_forecast_accuweather"
)

func main() {
	traffic.Conditions()
	weather.Forecast()
}
