// Package weather provides tools to help with weather forecast.
package weather

var (
    // CurrentCondition provides current weather condition, example includes rain, snow etc as a string. 
	CurrentCondition string
    // CurrentLocation provides the current location to be forecasted as a string.
	CurrentLocation  string
)

// Forecast returns a string value containing the CurrentLocation and the CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
