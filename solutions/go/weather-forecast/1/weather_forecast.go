// Package weather provides tools to track and forecast the current weather
// conditions for cities in Goblinocus.
package weather

var (
    // CurrentCondition stores the current weather condition for a city.
	CurrentCondition string

    // CurrentLocation stores the name of the city for which the weather
	// condition is currently being tracked.
	CurrentLocation  string
)

// Forecast sets the current weather condition and location,
// then returns a formatted weather forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
