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
			to := conv.FarhenheitToCelsius(from)
			fmt.Printf("%f°F er %f°C \n", from, to)
		case outArg == "K":
			to := conv.FarhenheitToKelvin(from)
			fmt.Printf("%f°F er %fK \n", from, to)
		case outArg == "F":
			to := fahrArg
			fmt.Printf("%f°F er %f°F \n", from, to)
		default:
			fmt.Printf("Passer gyldige arguenter og flagg \n")
		}

	case isFlagPassed("C"):
		from := celArg
		switch isFlagPassed("out") {
		case outArg == "C":
			to := celArg
			fmt.Printf("%f°C er %f°C \n", from, to)
		case outArg == "K":
			to := conv.CelciusToKelvin(from)
			fmt.Printf("%f°C er %fK \n", from, to)
		case outArg == "F":
			to := conv.CelciusToFarenheit(from)
			fmt.Printf("%f°C er %f°F \n", from, to)
		default:
			fmt.Printf("Passer gyldige arguenter og flagg \n")
		}

	case isFlagPassed("K"):
		from := kelvArg
		switch isFlagPassed("out") {
		case outArg == "C":
			to := conv.KelvinToCelcius(from)
			fmt.Printf("%fK er %f°C \n", from, to)
		case outArg == "K":
			to := kelvArg
			fmt.Printf("%fK er %fK \n", from, to)
		case outArg == "F":
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
