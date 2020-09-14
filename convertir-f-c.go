package main

import "fmt"

func main() {
	var tempF float64

	fmt.Print("Convertir de grados Fahrenheit a Celsius\n Ingresa la temperatura en Fahrenheit: ")
	fmt.Scan(&tempF)
	fmt.Printf("%f grados Fahrenheit son %f grados Celsius \n", tempF, (tempF-32)*5/9)
}
