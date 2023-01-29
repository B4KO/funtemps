package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
	"os"
	"strings"
)

var fahr float64
var cel float64
var kelv float64
var out string
var funfactArg string
var t string

func jannisTest() {
	fmt.Println(fahr, out, funfactArg)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println("0°F er -17.78°C")
	}
}

func setToUpper(out string, funfact string, t string) {
	out = strings.ToUpper(out)
	funfact = strings.ToUpper(funfact)
	t = strings.ToUpper(t)
}

// Validerer flaggene
func validateFlags(arguments []string) bool {

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

func init() {

	//os.Setenv("GO111MODULE ", "on")

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfactArg, "funfacts", "", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "", "\"fun-facts\" sin temperatur")
}

func main() {

	flag.Parse()

	//Normaliserer output slik at folk kan passsere små og store bokstaver uten feil
	setToUpper(out, funfactArg, t)

	//Sender inn alle argumentene passert utenom første som er program navnet
	if !validateFlags(os.Args[1:]) {
		//Slutter programmet om feil blir inputtet
		os.Exit(0)
	}

	switch {

	case isFlagPassed("F"):
		from := fahr
		switch isFlagPassed("out") {
		case out == "C":
			to := conv.FarhenheitToCelsius(from)
			fmt.Printf("%f°F er %f°C \n", from, to)
		case out == "K":
			to := conv.FarhenheitToKelvin(from)
			fmt.Printf("%f°F er %fK \n", from, to)
		case out == "F":
			to := fahr
			fmt.Printf("%f°F er %f°F \n", from, to)
		default:
			fmt.Printf("Passer gyldige arguenter og flagg \n")
		}

	case isFlagPassed("C"):
		from := cel
		switch isFlagPassed("out") {
		case out == "C":
			to := cel
			fmt.Printf("%f°C er %f°C \n", from, to)
		case out == "K":
			to := conv.CelciusToKelvin(from)
			fmt.Printf("%f°C er %fK \n", from, to)
		case out == "F":
			to := conv.CelciusToFarenheit(from)
			fmt.Printf("%f°C er %f°F \n", from, to)
		default:
			fmt.Printf("Passer gyldige arguenter og flagg \n")
		}

	case isFlagPassed("K"):
		from := kelv
		switch isFlagPassed("out") {
		case out == "C":
			to := conv.KelvinToCelcius(from)
			fmt.Printf("%fK er %f°C \n", from, to)
		case out == "K":
			to := kelv
			fmt.Printf("%fK er %fK \n", from, to)
		case out == "F":
			to := conv.KelvinToFarhenheit(from)
			fmt.Printf("%fK er %f°F \n", from, to)
		default:
			fmt.Printf("Passer gyldige arguenter og flagg \n")
		}

	case isFlagPassed("funfact"):
		case funfactArg == "SUN":




	default:
		fmt.Printf("Passer gyldige arguenter og flagg \n")


	}

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

switch isFlagPassed("t") {
case t == "C":
funfacts.GetFunFacts("SUN")
case t == "F":
funfacts.GetFunFacts("LUNA")
case t == "K":
funfacts.GetFunFacts("TERRA")
default:
fmt.Println("Passer gyldige argumenter og flagg \n")
}