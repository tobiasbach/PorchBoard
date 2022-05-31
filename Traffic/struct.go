package main

import "time"

type Traffic struct {
	FormatVersion string `json:"formatVersion"`
	Routes        []struct {
		Summary struct {
			LengthInMeters        int       `json:"lengthInMeters"`
			TravelTimeInSeconds   int       `json:"travelTimeInSeconds"`
			TrafficDelayInSeconds int       `json:"trafficDelayInSeconds"`
			TrafficLengthInMeters int       `json:"trafficLengthInMeters"`
			DepartureTime         time.Time `json:"departureTime"`
			ArrivalTime           time.Time `json:"arrivalTime"`
		} `json:"summary"`
		Legs []struct {
			Summary struct {
				LengthInMeters        int       `json:"lengthInMeters"`
				TravelTimeInSeconds   int       `json:"travelTimeInSeconds"`
				TrafficDelayInSeconds int       `json:"trafficDelayInSeconds"`
				TrafficLengthInMeters int       `json:"trafficLengthInMeters"`
				DepartureTime         time.Time `json:"departureTime"`
				ArrivalTime           time.Time `json:"arrivalTime"`
			} `json:"summary"`
			Points []struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"points"`
		} `json:"legs"`
		Sections []struct {
			StartPointIndex int    `json:"startPointIndex"`
			EndPointIndex   int    `json:"endPointIndex"`
			SectionType     string `json:"sectionType"`
			TravelMode      string `json:"travelMode"`
		} `json:"sections"`
	} `json:"routes"`
}
