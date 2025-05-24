package main

import (
	"fmt"
	"math"
)

func main() {

	var lado1, lado2 float64

	fmt.Print("Ingrese datos de lado1: ")
	fmt.Scanln(&lado1)

	fmt.Print("Ingrese datos de lado2: ")
	fmt.Scanln(&lado2)

	area := (lado1 * lado2) / 2
	fmt.Printf("El área es: %.2f \n", area)

	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa
	fmt.Printf("El perímetro es: %.2f", perimetro)

	gays := "\nAndy y Euge"
	fmt.Println(gays)
}
