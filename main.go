package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cels float64
var kelv float64
var out string
var funfact string
var temp string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfact, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&temp, "temp", "C", "bestemme hvilken temperatur som vises")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {
	flag.Parse()

	/*
		    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
		    pakkene implementeres.

		    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
		    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
		    hvor mange flagg og argumenter er spesifisert på kommandolinje.

		        fmt.Println("len(flag.Args())", len(flag.Args()))
				    fmt.Println("flag.NFlag()", flag.NFlag())

		    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
		    brukes for å utelukke ugyldige kombinasjoner:
		    -F, -C, -K kan ikke brukes samtidig
		    disse tre kan brukes med -out, men ikke med -funfacts
		    -funfacts kan brukes kun med -t
		    ...
		    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
		    implementert i flag-pakken og at den vil skrive ut "Usage" med
		    beskrivelsene av flagg-variablene, som angitt i parameter fire til
		    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	/*
		fmt.Println(fahr, out, funfact)

		fmt.Println("len(flag.Args())", len(flag.Args()))
		fmt.Println("flag.NFlag()", flag.NFlag())

		fmt.Println(isFlagPassed("out"))
	*/

	if isFlagPassed("F") {
		if out == "C" {
			fmt.Printf("%.2f°F er %.2f°C", fahr, conv.FarhenheitToCelsius(fahr))
		} else if out == "K" {
			fmt.Printf("%.2f°F er %.2f°K", fahr, conv.FarhenheitToKelvin(fahr))
		}
	} else if isFlagPassed("C") {
		if out == "F" {
			fmt.Printf("%.2f°C er %.2f°F", cels, conv.CelsiusToFahrenheit(cels))
		} else if out == "K" {
			fmt.Printf("%.2f°C er %.2f°K", cels, conv.CelsiusToKelvin(cels))
		}
	} else if isFlagPassed("K") {
		if out == "C" {
			fmt.Printf("%.2f°K er %.2f°C", kelv, conv.KelvinToCelsius(kelv))
		} else if out == "F" {
			fmt.Printf("%.2f°K er %.2f°F", kelv, conv.KelvinToFahrenheit(kelv))
		}
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
