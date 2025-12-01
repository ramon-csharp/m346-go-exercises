package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	fmt.Println("Celsius to Fahrenheit:")
	fmt.Printf("0°C = %.2f°F\n", convertCelsiusToFahrenheit(0))
	fmt.Printf("25°C = %.2f°F\n", convertCelsiusToFahrenheit(25))
	fmt.Printf("100°C = %.2f°F\n", convertCelsiusToFahrenheit(100))
	
	// TODO: call the function convertFahrenheitToCelsius
	fmt.Println("\nFahrenheit to Celsius:")
	fmt.Printf("32°F = %.2f°C\n", convertFahrenheitToCelsius(32))
	fmt.Printf("77°F = %.2f°C\n", convertFahrenheitToCelsius(77))
	fmt.Printf("212°F = %.2f°C\n", convertFahrenheitToCelsius(212))
}
