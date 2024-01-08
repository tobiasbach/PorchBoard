package main

import (
	// weather "PorchBoard/integrations/weather_forecast_accuweather"

	// car "PorchBoard/integrations/auxiliary_heating_car"
	temperature "PorchBoard/integrations/outdoor_temperature_viessmann"
	vbb "PorchBoard/integrations/public_transportation_vbb"
	traffic "PorchBoard/integrations/traffic_conditions_tomtom"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/board", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"trafficConditions": traffic.Conditions(),
			// "weatherForecast":    weather.Forecast(),
			"outsideTemperature": temperature.Outside(),
			"departures":         vbb.Departures(),
			// "family car":         car.WarmUpEngineAndCabin(),
		})
	})
	r.Run()
}
