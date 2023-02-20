package conv

import (
	"fmt"
	"strings"
)

// En funskjon som runder flyttall med en ubestemt mengde desimaler til et flyttall med to desimaler
/*func FloatRound(value float64) float64 {
	return math.Round(value*100) / 100
}*/

func FloatRound(value float64) string {
	str := fmt.Sprintf("%.2f", value)
	return strings.TrimRight(strings.TrimRight(str, "0"), ".")
}

// Konverteringsfunksjoner:

func FarhenheitToCelsius(fahrenheit float64) string {
	celsius := (fahrenheit - 32) * (5.0 / 9.0)
	return FloatRound(celsius)
}

func FarhenheitToKelvin(fahrenheit float64) string {
	kelvin := (fahrenheit-32)*(5.0/9) + 273.15
	return FloatRound(kelvin)
}

func CelsiusToFahrenheit(celsius float64) string {
	fahrenheit := celsius*(9/5.0) + 32
	return FloatRound(fahrenheit)
}

func CelsiusToKelvin(celsius float64) string {
	kelvin := celsius + 273.15
	return FloatRound(kelvin)
}

func KelvinToFahrenheit(kelvin float64) string {
	fahrenheit := (kelvin-273.15)*(9/5.0) + 32
	return FloatRound(fahrenheit)
}

func KelvinToCelsius(kelvin float64) string {
	celsius := kelvin - 273.15
	return FloatRound(celsius)
}
