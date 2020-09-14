package main

import "fmt"

func main() {
	var lado int

	fmt.Print("Calcular el área de un cuadrado\n Ingresa el tamaño de un lado del cuadrado: ")
	fmt.Scanf("%d", &lado)
	fmt.Printf("El area del cuadrado es: %d \n", lado*lado)
}
