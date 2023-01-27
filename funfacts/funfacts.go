package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
*/

func GetFunFacts(about string) []string {

	funfacts := makeFunFacts()

	switch about {
	case "SUN":
		return funfacts.Sun[:2]
	case "LUNA":
		return funfacts.Luna[:2]
	case "TERRA":
		return funfacts.Terra[:2]
	}
	return nil
}

func makeFunFacts() *FunFacts {

	funfacts := new(FunFacts)
	funfacts.Terra[0] = "Høyeste temperatur målt på Jordens overflate"
	funfacts.Terra[2] = "Laveste temperatur målt på Jordens overflate"
	funfacts.Terra[3] = "Temperatur i Jordens indre kjerne"
	funfacts.Sun[0] = "Temperatur i Solens kjerne"
	funfacts.Sun[1] = "Temperatur på ytre lag av Solen"
	funfacts.Luna[0] = "Temperatur på Månens overflate om natten"
	funfacts.Luna[1] = "Temperatur på Månens overflate om dagen"

	return funfacts
}

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}
