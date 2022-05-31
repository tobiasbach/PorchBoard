package main

import "time"

type Weather struct {
	DateTime         time.Time `json:"DateTime"`
	EpochDateTime    int       `json:"EpochDateTime"`
	WeatherIcon      int       `json:"WeatherIcon"`
	IconPhrase       string    `json:"IconPhrase"`
	HasPrecipitation bool      `json:"HasPrecipitation"`
	IsDaylight       bool      `json:"IsDaylight"`
	Temperature      struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"Temperature"`
	RealFeelTemperature struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
		Phrase   string  `json:"Phrase"`
	} `json:"RealFeelTemperature"`
	RealFeelTemperatureShade struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
		Phrase   string  `json:"Phrase"`
	} `json:"RealFeelTemperatureShade"`
	WetBulbTemperature struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"WetBulbTemperature"`
	DewPoint struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"DewPoint"`
	Wind struct {
		Speed struct {
			Value    float64 `json:"Value"`
			Unit     string  `json:"Unit"`
			UnitType int     `json:"UnitType"`
		} `json:"Speed"`
		Direction struct {
			Degrees   int    `json:"Degrees"`
			Localized string `json:"Localized"`
			English   string `json:"English"`
		} `json:"Direction"`
	} `json:"Wind"`
	WindGust struct {
		Speed struct {
			Value    float64 `json:"Value"`
			Unit     string  `json:"Unit"`
			UnitType int     `json:"UnitType"`
		} `json:"Speed"`
		Direction struct {
			Degrees   int    `json:"Degrees"`
			Localized string `json:"Localized"`
			English   string `json:"English"`
		} `json:"Direction"`
	} `json:"WindGust"`
	RelativeHumidity       int `json:"RelativeHumidity"`
	IndoorRelativeHumidity int `json:"IndoorRelativeHumidity"`
	Visibility             struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"Visibility"`
	Ceiling struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"Ceiling"`
	UVIndex                  int    `json:"UVIndex"`
	UVIndexText              string `json:"UVIndexText"`
	PrecipitationProbability int    `json:"PrecipitationProbability"`
	ThunderstormProbability  int    `json:"ThunderstormProbability"`
	RainProbability          int    `json:"RainProbability"`
	SnowProbability          int    `json:"SnowProbability"`
	IceProbability           int    `json:"IceProbability"`
	TotalLiquid              struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"TotalLiquid"`
	Rain struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"Rain"`
	Snow struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"Snow"`
	Ice struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"Ice"`
	CloudCover         int `json:"CloudCover"`
	Evapotranspiration struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"Evapotranspiration"`
	SolarIrradiance struct {
		Value    float64 `json:"Value"`
		Unit     string  `json:"Unit"`
		UnitType int     `json:"UnitType"`
	} `json:"SolarIrradiance"`
	MobileLink string `json:"MobileLink"`
	Link       string `json:"Link"`
}
