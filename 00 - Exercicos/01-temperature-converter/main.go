package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 7) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	// Variáveis com valores iniciais
	celsius := 36.0
	fahrenheit := 98.6

	// Chamando funções e armazenando resultados
	resultF := celsiusToFahrenheit(celsius)
	resultC := fahrenheitToCelsius(fahrenheit)

	// Imprimindo resultados formatados
	fmt.Printf("%.2f°C equivale a %.2f°F\n", celsius, resultF)
	fmt.Printf("%.2f°F equivale a %.2f°C\n", fahrenheit, resultC)
}