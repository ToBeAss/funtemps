package funfacts

// Definerer en klasse: FunFacts
type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string) []string {

	var funfact FunFacts // Lager en instans av klassen/typen FunFacts

	// Tilegner verdier
	funfact.Sun = []string{"Temperatur i Solens kjerne", "Temperatur på ytre lag av Solen"}
	funfact.Luna = []string{"Temperatur på Månens overflate om natten", "Temperatur på Månens overflate om dagen"}
	funfact.Terra = []string{"Høyeste temperatur målt på Jordens overflate", "Laveste temperatur målt på Jordens overflate", "Temperatur i Jordens indre kjerne"}

	// Logikk for utvelging av funfact basert på input: about
	if about == "sun" {
		return funfact.Sun
	} else if about == "luna" {
		return funfact.Luna
	} else if about == "terra" {
		return funfact.Terra
	}
	return nil
}
