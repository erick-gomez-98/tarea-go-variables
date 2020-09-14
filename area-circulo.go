package main

import (
	"fmt"
	"math"
)

func main() {
	var radio float64

	fmt.Print("Calcular el área de un círculo\n Ingresa la medida del radio: ")
	fmt.Scan(&radio)
	fmt.Printf("El area del circulo es: %f \n", math.Pi*(math.Pow(radio, 2)))
}
