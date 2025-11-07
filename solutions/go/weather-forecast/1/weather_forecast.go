// Package weather includes functions that provide forecast info.
package weather

var (
    // CurrentCondition stores current weather conditions.
	CurrentCondition string
    // CurrentLocation stores current location.
	CurrentLocation  string
)

// Forecast provides an overview of the location and current weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
