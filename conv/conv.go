package conv

import (
	/*"fmt"
	"strings"*/
)

// En funskjon som runder flyttall med en ubestemt mengde desimaler til et flyttall med to desimaler
/*
	func FloatRound(value float64) float64 {
		return math.Round(value*100) / 100
	}
*/

// Lignende funksjon som den over, men her blir unødvendige nuller (0) tatt bort
/*
	func FloatRound(value float64) string {
		str := fmt.Sprintf("%.2f", value)
		return strings.TrimRight(strings.TrimRight(str, "0"), ".")
	}
*/

// Konverteringsfunksjoner:

func FarhenheitToCelsius(fahrenheit float64) float64 {
	celsius := (fahrenheit - 32) * (5.0 / 9.0)
	return celsius
}

func FarhenheitToKelvin(fahrenheit float64) float64 {
	kelvin := (fahrenheit-32)*(5.0/9) + 273.15
	return kelvin
}

func CelsiusToFahrenheit(celsius float64) float64 {
	fahrenheit := celsius*(9/5.0) + 32
	return fahrenheit
}

func CelsiusToKelvin(celsius float64) float64 {
	kelvin := celsius + 273.15
	return kelvin
}

func KelvinToFahrenheit(kelvin float64) float64 {
	fahrenheit := (kelvin-273.15)*(9/5.0) + 32
	return fahrenheit
}

func KelvinToCelsius(kelvin float64) float64 {
	celsius := kelvin - 273.15
	return celsius
}
