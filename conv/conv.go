package conv

import "github.com/B4KO/funtemps/functions"

func FarhenheitToCelsius(value float64) float64 {
	return functions.RoundFloat((value-32.0)*(5.0/9.0), 2)
}

func CelciusToFarenheit(value float64) float64 {
	return functions.RoundFloat(value*(9.0/5.0)+32.0, 2)
}

func CelciusToKelvin(value float64) float64 {
	return functions.RoundFloat(value+273.15, 2)
}
func KelvinToCelcius(value float64) float64 {
	return functions.RoundFloat(value-273.15, 2)
}

func KelvinToFarhenheit(value float64) float64 {
	return functions.RoundFloat((value-273.15)*(9.0/5.0)+32.0, 2)
}

func FarhenheitToKelvin(value float64) float64 {
	return functions.RoundFloat((value-32.0)*(5.0/9.0)+273.15, 2)
}
