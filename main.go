package main

import (
	// weather "PorchBoard/api_integrations/weather_forecast_accuweather"

	temperature "PorchBoard/api_integrations/outdoor_temperature_viessmann"
	vbb "PorchBoard/api_integrations/public_transportation_vbb"
	traffic "PorchBoard/api_integrations/traffic_conditions_tomtom"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/board", func(c *gin.Context) {
		trafficConditionsValue := traffic.Conditions()
		// weatherForecastValue := weather.Forecast()
		outsideTemperatureValue := temperature.Outside()
		departuresValue := vbb.Departures()
		c.JSON(200, gin.H{
			"trafficConditions": trafficConditionsValue,
			// "weatherForecast": weatherForecastValue,
			"outsideTemperature": outsideTemperatureValue,
			"departures":         departuresValue,
		})
	})
	r.Run()
}

// until a specific port is defined, serving on localhost:8080/board
