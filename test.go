package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	const (
		Lunes = iota + 1
		Martes
		Miercoles
		Jueves
		Viernes
		Sabado
		Domingo
	)
	fmt.Println(Viernes) //Imprime 5

	const Pi float32 = 3.14 //Deben declararse e inicializarse
	const Euler = 2.71828   //No es necesario declarar el tipo de datos
	//Se recomienda crearlos con mayúscula para para que puedan ser leídos por cualquier paquete

	const (
		avogadro = 606562
		x        = 100
		y        = 0b1010 //binarios
		z        = 0o12   //octal
		w        = 0xFF   //Hexadecimal
	)

	fmt.Println(avogadro, x, y, z, w)

	var nombresPeliculas string = "Peliculas Cat"
	fmt.Println(nombresPeliculas, x, y, 12)

	var firstName1 string                               //Declarar una sola variable
	var middleName1, lastName1 string = "Iván", "Jalid" //Declarar varias variables del mismo tipo
	var age1 int

	firstName1 = "Agus"
	age1 = 89

	fmt.Println(firstName1, middleName1, lastName1, age1)

	var ( //Declarar todas las variables juntas
		firstName2, lastName2 string
		age2                         = 89 //Cuando una variable se crea e inicializa no es necesario colocar el tipo de dato
		dni2                  string = "34004888"
	)

	listaDePutos := [...]string{"Germán Morini", "Andrés Vergachica"}

	fmt.Println("En la lista de putos solo está", listaDePutos[0], "y", listaDePutos[1])
	fmt.Println(firstName2, lastName2, age2, dni2)

	//Se pueden crear todas las variables en la misma línea
	var firstName3, lastName3, age3 = "Agus", "Alici", 89

	fmt.Println(firstName3, lastName3, age3)

	dni1 := "34004888"                                 //Declarar una variable e inicializarla
	firstName4, lastName4, age4 := "Agus", "Alici", 89 //Declarar e inicializar (forma más usada DENTRO de las funciones)
	firstName4 = "Agustina"                            //Modificar una variable
	fmt.Println(dni1, firstName4, lastName4, age4)
}
