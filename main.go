package main

import (
	"flag"
	"fmt"
	"funtemps/conv"
	"funtemps/functions"
	"os"
)

var fahrArg float64
var celArg float64
var kelvArg float64
var outArg string
var funfactArg string
var t string

func init() {
	flag.Float64Var(&fahrArg, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celArg, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvArg, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&outArg, "out", "", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfactArg, "funfacts", "", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "", "\"fun-facts\" sin temperatur")
}

func main() {

	flag.Parse()

	//Normaliserer output slik at folk kan passsere små og store bokstaver uten å få feil
	functions.SetToUpper(outArg, funfactArg, t)

	//Sender inn alle argumentene passert, utenom program navnet
	if !functions.ValidateFlags(os.Args[1:]) {
		//Avslutter programmet
		os.Exit(0)
	}

	switch {

	case isFlagPassed("F"):
		from := fahrArg
		switch isFlagPassed("out") {
		case outArg == "C":
			to := functions.FormatFloat(conv.FarhenheitToCelsius(from))
			fmt.Printf("%s°F er %s°C \n", os.Args[2], to)
		case outArg == "K":
			to := functions.FormatFloat(conv.FarhenheitToKelvin(from))
			fmt.Printf("%s°F er %sK \n", os.Args[2], to)
		case outArg == "F":
			to := functions.FormatFloat(fahrArg)
			fmt.Printf("%s°F er %s°F \n", os.Args[2], to)
		default:
			fmt.Printf("Passer gyldige arguenter og flagg \n")
			os.Exit(0)
		}

	case isFlagPassed("C"):
		from := celArg
		switch isFlagPassed("out") {
		case outArg == "C":
			to := functions.FormatFloat(celArg)
			fmt.Printf("%s°C er %s°C \n", os.Args[2], to)
		case outArg == "K":
			to := functions.FormatFloat(conv.CelciusToKelvin(from))
			fmt.Printf("%s°C er %sK \n", os.Args[2], to)
		case outArg == "F":
			to := functions.FormatFloat(conv.CelciusToFarenheit(from))
			fmt.Printf("%s°C er %s°F \n", os.Args[2], to)
		default:
			fmt.Printf("Passer gyldige arguenter og flagg \n")
			os.Exit(0)
		}

	case isFlagPassed("K"):
		from := kelvArg
		switch isFlagPassed("out") {
		case outArg == "C":
			to := functions.FormatFloat(conv.KelvinToCelcius(from))
			fmt.Printf("%sK er %s°C \n", os.Args[2], to)
		case outArg == "K":
			to := functions.FormatFloat(kelvArg)
			fmt.Printf("%sK er %sK \n", os.Args[2], to)
		case outArg == "F":
			to := functions.FormatFloat(conv.KelvinToFarhenheit(from))
			fmt.Printf("%sK er %s°F \n", os.Args[2], to)
		default:
			fmt.Printf("Passer gyldige arguenter og flagg \n")
			os.Exit(0)
		}

		//case isFlagPassed("funfact"):
		//case funfactArg == "SUN":
		//
		//default:
		//	fmt.Printf("Passer gyldige arguenter og flagg \n")

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

//switch isFlagPassed("t") {
//case t == "C":
//funfacts.GetFunFacts("SUN")
//case t == "F":
//funfacts.GetFunFacts("LUNA")
//case t == "K":
//funfacts.GetFunFacts("TERRA")
//default:
//fmt.Println("Passer gyldige argumenter og flagg \n")
//}
