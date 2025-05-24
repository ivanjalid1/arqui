package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completada  bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
	fmt.Println("\nTarea agregada correctamente")
}

func (l *ListaTareas) mostrarTareas() {

	fmt.Println("Listado de tareas")
	fmt.Println("\n=================")

	for index, t := range l.tareas {
		fmt.Printf("\n%d. \n - %s - %s - %t", index+1, t.nombre, t.descripcion, t.completada)

	}
	fmt.Println("\n=================")
}

func (l *ListaTareas) marcarCompletado(indice int) {
	l.tareas[indice].completada = true
}

func (l *ListaTareas) eliminarTarea(indice int) {
	l.tareas = append(l.tareas[:indice], l.tareas[indice+1:]...)
}

func (l *ListaTareas) editarTarea(indice int, t Tarea) {
	l.tareas[indice] = t
}

func displayMenu() int {
	fmt.Println("Selecciona una opción:" +
		"\n1. Agregar tarea" +
		"\n2. Marcar tarea como completada" +
		"\n3. Editar Tarea" +
		"\n4. Eliminar Tarea" +
		"\n5. Salir",
	)
	var opcion int
	fmt.Scanln(&opcion)
	return opcion
}

func crearTarea(action string) Tarea {
	fmt.Print("Ingrese el nombre de la tarea %s: \n", action)
	leer := bufio.NewReader(os.Stdin)
	nom, _ := leer.ReadString('\n')
	fmt.Print("Ingrese la descripción de la tarea %s: \n", action)
	desc, _ := leer.ReadString('\n')
	return Tarea{nombre: nom, descripcion: desc}
}

func (l *ListaTareas) obtenerIndice(action string) (int, error) {
	fmt.Printf("Ingresar indice de la tarea %s: ", action)
	var indice int
	fmt.Scanln(&indice)

	if indice < 0 || indice >= len(l.tareas)() {
		return -1, errors.New("Valor de índice correcto")
	}
	return indice, nil

}

func main() {
	lista := ListaTareas{}
	for {
		opcion := displayMenu()
		lista.mostrarTareas()
		switch opcion {
		case 1:
			t := crearTarea("agregar")
			lista.agregarTarea(t)

		case 2:
			indice, err := lista.obtenerIndice("marcar como completada")

			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			lista.marcarCompletado(indice)
		case 3:
			indice, err := lista.obtenerIndice("Editar")

			if err != nil {
				fmt.Println("Error: ", err)
				break
			}

			t := crearTarea("editar")
			lista.editarTarea(indice, t)
		case 4:
			indice, err := lista.obtenerIndice("eliminar")

			if err != nil {
				fmt.Println("Error: ", err)
				break
			}

			lista.eliminarTarea(indice)
		case 5:
			fmt.Println("Adios!!")

		default:
			fmt.Println("Error, ingrese un número correcto")
		}
		lista.mostrarTareas()
	}
}
