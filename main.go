package main

import (
	temperature "PorchBoard/api_integrations/outdoor_temperature_viessmann"
	vbb "PorchBoard/api_integrations/public_transportation_vbb"
	traffic "PorchBoard/api_integrations/traffic_conditions_tomtom"
	weather "PorchBoard/api_integrations/weather_forecast_accuweather"
)

func main() {
	traffic.Conditions()
	weather.Forecast()
	temperature.Outside()
	vbb.Departures()
}
