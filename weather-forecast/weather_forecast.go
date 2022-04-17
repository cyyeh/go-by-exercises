// Package weather provides tools to forecast weather
// given city location and current condition.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current city location.
var CurrentLocation string

// Forecast returns a string value that describes the forecasting result
// given city and condition strings.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
