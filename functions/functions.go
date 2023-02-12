package functions

import (
	"flag"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"strings"
)

func FormatFloat(val float64) string {
	num := val

	if math.Mod(num, 1) == 0.00 || math.Mod(num, 1) == -0.00 {
		num := int(num)
		p := message.NewPrinter(language.English)
		s := strings.Replace(p.Sprintf("%.2d", num), ",", " ", -1)

		return s
	}

	p := message.NewPrinter(language.English)
	s := strings.Replace(p.Sprintf("%.2f", num), ",", " ", -1)

	return s

}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func SetToUpper(out string, funfact string, t string) {
	out = strings.ToUpper(out)
	funfact = strings.ToUpper(funfact)
	t = strings.ToUpper(t)
}

// Validerer flaggene
func ValidateFlags(arguments []string) bool {

	//Ser om riktig antall flag og argumenter blir passert
	if flag.NFlag() != 2 {
		return false
	}
	if len(flag.Args()) != 2 {
		return false
	}

	//Deffinerer de riktige flaggene i første og andre slot
	FirstFlags := []string{"-F", "-C", "-K", "-funfacts"}
	SecondFlags := []string{"-out", "-t"}

	//Validerer om riktig kombinasjon av argumenter blir passert
	if arguments[0] == FirstFlags[0] || arguments[0] == FirstFlags[1] || arguments[0] == FirstFlags[2] {
		if arguments[2] == SecondFlags[0] {
			return true
		}
	} else if arguments[1] == FirstFlags[3] {
		if arguments[2] == SecondFlags[1] {
			return true
		}
	}
	return false
}
