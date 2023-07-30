// Package weather_forecast provides tools for forecasting weather.
package weather_forecast

// CurrentCondition describes the current weather condition (e.g. "rainy").
var CurrentCondition string

// CurrentLocation describes the current location (e.g. "Switzerland").
var CurrentLocation string

// Forecast returns a string that describes weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
