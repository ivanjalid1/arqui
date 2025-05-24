package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

func play() {
	numAleatorio := rand.Intn(100)
	var numIngresado int

	for i := 0; i < 10; i++ {
		fmt.Printf("Ingrese un número (intentos restantes %d): ", 10-i)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("Ganaste, lo hiciste en ", i+1, "intentos")
			displayMenu()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("No acertaste, el número es menor")
		} else {
			fmt.Println("No acertaste, el número es mayor")
		}
	}
	fmt.Println("No ganaste, el número era: ", numAleatorio)
	displayMenu()
}

func displayMenu() {
	fmt.Println("¿Desea jugar nuevamente? (s/n)")
	var opt string
	fmt.Scanln(&opt)

	switch opt {
	case "s":
		play()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		displayMenu()
	}

}
