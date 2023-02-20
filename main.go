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
	flag.Float64Var(&fahr, "F", 0.0, "viser temperatur i grader fahrenheit (må være tall)")
	flag.Float64Var(&cels, "C", 0.0, "viser temperatur i grader celsius (må være tall)")
	flag.Float64Var(&kelv, "K", 0.0, "viser temperatur i grader kelvin (må være tall)")
	flag.StringVar(&out, "out", "C", "konverterer temperatur til C, F eller K")

	flag.StringVar(&funfact, "funfacts", "sun", "funfacts om sun, luna og terra")
	flag.StringVar(&temp, "temp", "C", "velger hvilken temperatur som vises sammen med funfacts")
}

func main() {
	flag.Parse()

	if isFlagPassed("out") { // LOGIKK FOR KONVERTERING AV GRADER
		if isFlagPassed("F") { // Input
			if out == "C" {
				fmt.Printf("%s°F er %s°C\n", conv.FloatRound(fahr), conv.FarhenheitToCelsius(fahr))
				// %s viser Strings, som har vært gjennom funksjonen FloatRound i conv-pakken
			} else if out == "K" {
				fmt.Printf("%s°F er %s°K\n", conv.FloatRound(fahr), conv.FarhenheitToKelvin(fahr))
			} else {
				fmt.Println("Flag out has to be C or K")
			}
		} else if isFlagPassed("C") {
			if out == "F" {
				fmt.Printf("%s°C er %s°F\n", conv.FloatRound(cels), conv.CelsiusToFahrenheit(cels))
			} else if out == "K" {
				fmt.Printf("%s°C er %s°K\n", conv.FloatRound(cels), conv.CelsiusToKelvin(cels))
			} else {
				fmt.Println("Flag out has to be F or K")
			}
		} else if isFlagPassed("K") {
			if out == "C" {
				fmt.Printf("%s°K er %s°C\n", conv.FloatRound(kelv), conv.KelvinToCelsius(kelv))
			} else if out == "F" {
				fmt.Printf("%s°K er %s°F\n", conv.FloatRound(kelv), conv.KelvinToFahrenheit(kelv))
			} else {
				fmt.Println("Flag out has to be C or F")
			}
		}
	} else if isFlagPassed("funfacts") && isFlagPassed("temp") { // LOGIKK FOR FUNFACTS
		if temp == "C" {
			if funfact == "sun" {
				// funfact[0]: "Temperatur i Solens kjerne", funfact[1]: "Temperatur på ytre lag av Solen"
				fmt.Printf("%s er %s°C\n%s er %s°C", funfacts.GetFunFacts(funfact)[0], conv.FloatRound(15000000.0), funfacts.GetFunFacts(funfact)[1], conv.KelvinToCelsius(5778))
				// %s viser samme verdi som inputverdi (i dette tilfellet: String)
				// \n hopper ned en linje i terminalen og separerer dermed output-setningene
				// funfacts.GetFunFacts(funfact) henter et array/slice med funfacts. En enkelt funfact (en String) velges med index (f.eks. [0]) bak funksjonskallet
				// Verdien til %s blir puttet inn gjennom conv-pakken (conv.FloatRound(15000000.0) eller conv.KelvinToCelsius(5778))
			} else if funfact == "luna" {
				// funfact[0]: "Temperatur på Månens overflate om natten", funfact[1]: "Temperatur på Månens overflate om dagen"
				fmt.Printf("%s er %s°C\n%s er %s°C", funfacts.GetFunFacts(funfact)[0], conv.FloatRound(-183.0), funfacts.GetFunFacts(funfact)[1], conv.FloatRound(106.0))
			} else if funfact == "terra" {
				// funfact[0]: "Høyeste temperatur målt på Jordens overflate", funfact[1]: "Laveste temperatur målt på Jordens overflate", funfact[2]: "Temperatur i Jordens indre kjerne"
				fmt.Printf("%s er %s°C\n%s er %s°C\n%s er %s°C", funfacts.GetFunFacts(funfact)[0], conv.FloatRound(56.7), funfacts.GetFunFacts(funfact)[1], conv.FloatRound(-89.4), funfacts.GetFunFacts(funfact)[2], conv.KelvinToCelsius(9392))
			} else {
				fmt.Printf("%s is not defined, try sun, luna or terra", funfact)
			}
		} else if temp == "F" {
			if funfact == "sun" {
				fmt.Printf("%s er %s°F\n%s er %s°F", funfacts.GetFunFacts(funfact)[0], conv.CelsiusToFahrenheit(15000000), funfacts.GetFunFacts(funfact)[1], conv.KelvinToFahrenheit(5778))
			} else if funfact == "luna" {
				fmt.Printf("%s er %s°F\n%s er %s°F", funfacts.GetFunFacts(funfact)[0], conv.CelsiusToFahrenheit(-183), funfacts.GetFunFacts(funfact)[1], conv.CelsiusToFahrenheit(106))
			} else if funfact == "terra" {
				fmt.Printf("%s er %s°F\n%s er %s°F\n%s er %s°F", funfacts.GetFunFacts(funfact)[0], conv.FloatRound(134.0), funfacts.GetFunFacts(funfact)[1], conv.CelsiusToFahrenheit(-89.4), funfacts.GetFunFacts(funfact)[2], conv.KelvinToFahrenheit(9392))
			} else {
				fmt.Printf("%s is not defined, try sun, luna or terra", funfact)
			}
		} else if temp == "K" {
			if funfact == "sun" {
				fmt.Printf("%s er %s°K\n%s er %s°K", funfacts.GetFunFacts(funfact)[0], conv.CelsiusToKelvin(15000000), funfacts.GetFunFacts(funfact)[1], conv.FloatRound(5778.0))
			} else if funfact == "luna" {
				fmt.Printf("%s er %s°K\n%s er %s°K", funfacts.GetFunFacts(funfact)[0], conv.CelsiusToKelvin(-183), funfacts.GetFunFacts(funfact)[1], conv.CelsiusToKelvin(106))
			} else if funfact == "terra" {
				fmt.Printf("%s er %s°K\n%s er %s°K\n%s er %s°K", funfacts.GetFunFacts(funfact)[0], conv.FloatRound(329.82), funfacts.GetFunFacts(funfact)[1], conv.CelsiusToKelvin(-89.4), funfacts.GetFunFacts(funfact)[2], conv.FloatRound(9392.0))
			} else {
				fmt.Printf("%s is not defined, try sun, luna or terra", funfact)
			}
		} else {
			fmt.Println("Flag temp has to be C, F or K")
		}
	} else {
		fmt.Println("Missing flags")
		fmt.Println("C/F/K + out")
		fmt.Println("funfacts + temp")
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
