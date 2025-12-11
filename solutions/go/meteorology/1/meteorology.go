package meteorology

import (
	"fmt"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

func (tempUnit TemperatureUnit) String() string {
	if tempUnit == 0 {
		return "°C"
	}

	return "°F"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (temp Temperature) String() string {
	return fmt.Sprintf("%v %v", temp.degree, temp.unit.String())
}

// Add a String method to the Temperature type

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

func (sUnit SpeedUnit) String() string {
	if sUnit == 0 {
		return "km/h"
	}

	return "mph"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit.String())
}

// Add a String method to Speed

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (meteoData MeteorologyData) String() string {
	return fmt.Sprintf(
		"%v: %v, Wind %v at %v, %v%% Humidity",
		meteoData.location,
		meteoData.temperature,
		meteoData.windDirection,
		meteoData.windSpeed.String(),
		meteoData.humidity)
}
