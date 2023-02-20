package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
	"funtemps/funfacts"
)

// Definerer flag-variabler i hoved-"scope"
var fahr float64
var cels float64
var kelv float64
var out string
var funfact string
var temp string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene er initialisert.
func init() {
	// Definerer og initialiserer flagg-variablene
	// Det siste argumentet forklarer hva variabelen blir brukt til
	flag.Float64Var(&fahr, "F", 0.0, "viser temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "viser temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "viser temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "konverterer temperatur til C, F eller K")

	flag.StringVar(&funfact, "funfacts", "sun", "funfacts om sun, luna og terra")
	flag.StringVar(&temp, "temp", "C", "velger hvilken temperatur som vises sammen med funfacts")
}

func main() {
	flag.Parse()

	// Logikk for konvertering av grader
	if isFlagPassed("F") { // Input
		if out == "C" {
			fmt.Printf("%g°F er %g°C", fahr, conv.FarhenheitToCelsius(fahr))
			// %g viser flyttall, med kun "nødvendige" desimaler (3.00 blir 3, 3.10 blir 3.1, 3.14 blir 3.14)
		} else if out == "K" {
			fmt.Printf("%g°F er %g°K", fahr, conv.FarhenheitToKelvin(fahr))
		}
	} else if isFlagPassed("C") {
		if out == "F" {
			fmt.Printf("%g°C er %g°F", cels, conv.CelsiusToFahrenheit(cels))
		} else if out == "K" {
			fmt.Printf("%g°C er %g°K", cels, conv.CelsiusToKelvin(cels))
		}
	} else if isFlagPassed("K") {
		if out == "C" {
			fmt.Printf("%g°K er %g°C", kelv, conv.KelvinToCelsius(kelv))
		} else if out == "F" {
			fmt.Printf("%g°K er %g°F", kelv, conv.KelvinToFahrenheit(kelv))
		}
	}

	// Logikk for funfacts
	if temp == "C" {
		if funfact == "sun" {
			// funfact[0]: "Temperatur i Solens kjerne", funfact[1]: "Temperatur på ytre lag av Solen"
			fmt.Printf("%v er %g°C\n%v er %g°C", funfacts.GetFunFacts(funfact)[0], 15000000.0, funfacts.GetFunFacts(funfact)[1], conv.KelvinToCelsius(5778))
			// %v viser samme verdi som inputverdi (i dette tilfellet: String)
			// \n hopper ned en linje i terminalen og separerer dermed output-setningene
			// funfacts.GetFunFacts(funfact) henter et array/slice med funfacts. En enkelt funfact (en String) velges med index (f.eks. [0]) bak funksjonskallet
			// Verdien til %g blir puttet enten puttet inn direkte som flyttall (15000000.0), eller gjennom conv-pakken (conv.KelvinToCelsius(5778))
		} else if funfact == "luna" {
			// funfact[0]: "Temperatur på Månens overflate om natten", funfact[1]: "Temperatur på Månens overflate om dagen"
			fmt.Printf("%v er %g°C\n%v er %g°C", funfacts.GetFunFacts(funfact)[0], -183.0, funfacts.GetFunFacts(funfact)[1], 106.0)
		} else if funfact == "terra" {
			// funfact[0]: "Høyeste temperatur målt på Jordens overflate", funfact[1]: "Laveste temperatur målt på Jordens overflate", funfact[2]: "Temperatur i Jordens indre kjerne"
			fmt.Printf("%v er %g°C\n%v er %g°C\n%v er %g°C", funfacts.GetFunFacts(funfact)[0], 56.7, funfacts.GetFunFacts(funfact)[1], -89.4, funfacts.GetFunFacts(funfact)[2], conv.KelvinToCelsius(9392))
		}
	} else if temp == "F" {
		if funfact == "sun" {
			fmt.Printf("%v er %g°F\n%v er %g°F", funfacts.GetFunFacts(funfact)[0], conv.CelsiusToFahrenheit(15000000), funfacts.GetFunFacts(funfact)[1], conv.KelvinToFahrenheit(5778))
		} else if funfact == "luna" {
			fmt.Printf("%v er %g°F\n%v er %g°F", funfacts.GetFunFacts(funfact)[0], conv.CelsiusToFahrenheit(-183), funfacts.GetFunFacts(funfact)[1], conv.CelsiusToFahrenheit(106))
		} else if funfact == "terra" {
			fmt.Printf("%v er %g°F\n%v er %g°F\n%v er %g°F", funfacts.GetFunFacts(funfact)[0], 134.0, funfacts.GetFunFacts(funfact)[1], conv.CelsiusToFahrenheit(-89.4), funfacts.GetFunFacts(funfact)[2], conv.KelvinToFahrenheit(9392))
		}
	} else if temp == "K" {
		if funfact == "sun" {
			fmt.Printf("%v er %g°K\n%v er %g°K", funfacts.GetFunFacts(funfact)[0], conv.CelsiusToKelvin(15000000), funfacts.GetFunFacts(funfact)[1], 5778.0)
		} else if funfact == "luna" {
			fmt.Printf("%v er %g°K\n%v er %g°K", funfacts.GetFunFacts(funfact)[0], conv.CelsiusToKelvin(-183), funfacts.GetFunFacts(funfact)[1], conv.CelsiusToKelvin(106))
		} else if funfact == "terra" {
			fmt.Printf("%v er %g°K\n%v er %g°K\n%v er %g°K", funfacts.GetFunFacts(funfact)[0], 329.82, funfacts.GetFunFacts(funfact)[1], conv.CelsiusToKelvin(-89.4), funfacts.GetFunFacts(funfact)[2], 9392.0)
		}
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
