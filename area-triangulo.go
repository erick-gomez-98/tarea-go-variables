package main

import "fmt"

func main() {
	var base, altura float64

	fmt.Println("Calcular el área de un triangulo")
	fmt.Println("Ingresa la base: ")
	fmt.Scanf("%f", &base)
	fmt.Println("Ingresa la altura: ")
	fmt.Scanf("%f", &altura)
	fmt.Printf("El area del triangulo es: %f \n", (base*altura)/2)
}
