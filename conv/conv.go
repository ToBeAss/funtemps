package conv

import (
	"math"
)

/*
I denne pakken skal alle konverteringfunksjonene
implementeres. Bruk engelsk.

	FarhenheitToCelsius
	CelsiusToFahrenheit
	KelvinToFarhenheit
	...
*/
func FloatRound(value float64) float64 {
	return math.Round(value*100) / 100
}

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(fahrenheit float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	celsius := (fahrenheit - 32) * (5.0 / 9)
	return FloatRound(celsius) // Skal bli 56.67
}

// De andre konverteringsfunksjonene implementere her
// ...
func FarhenheitToKelvin(fahrenheit float64) float64 {
	kelvin := (fahrenheit-32)*(5.0/9) + 273.15
	return FloatRound(kelvin)
}

func CelsiusToFahrenheit(celsius float64) float64 {
	fahrenheit := celsius*(9/5.0) + 32
	return FloatRound(fahrenheit)
}

func CelsiusToKelvin(celsius float64) float64 {
	kelvin := celsius + 273.15
	return FloatRound(kelvin)
}

func KelvinToFahrenheit(kelvin float64) float64 {
	fahrenheit := (kelvin-273.15)*(9/5.0) + 32
	return FloatRound(fahrenheit)
}

func KelvinToCelsius(kelvin float64) float64 {
	celsius := kelvin - 273.15
	return FloatRound(celsius)
}
