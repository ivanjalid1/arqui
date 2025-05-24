package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const fileName = "lista-contactos.json"

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"` // la letra es mayuscula pq queremos que sea una variable publica
}

func saveContactToFile(contactList []Contact) error {
	file, err := os.Create(fileName)

	if err != nil {
		return errors.New("No se pudo crear el archivo")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contactList)

	if err != nil {
		return errors.New("No se pudo formatear el contacto")
	}

	return nil
}

func loadContactList(contactList *[]Contact) error {
	file, err := os.Open(fileName)

	if err != nil {
		return errors.New("No se pudo abrir el archivo")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(contactList)

	if err != nil {
		return errors.New("Error al decodear los contactos")
	}

	return nil
}

func main() {
	var contactList []Contact

	err := loadContactList(&contactList)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for {
		fmt.Print("==== GESTOR DE CONTACTOS ====" +
			"\n1. Agregar un contacto" +
			"\n2. Mostrar todos los contactos" +
			"\n3. Salir" +
			"\n Elige un opción\n",
		)

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			//agregar
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Nombre: ")
			nombre, _ := reader.ReadString('\n')

			fmt.Println("Email ")
			email, _ := reader.ReadString('\n')

			fmt.Println("Telefono: ")
			telefono, _ := reader.ReadString('\n')

			nombre = strings.TrimSpace(nombre) // limpia espacios en blanco
			email = strings.TrimSpace(email)
			telefono = strings.TrimSpace(telefono)

			con := Contact{Name: nombre, Email: email, Phone: telefono}
			contactList = append(contactList, con)

			err = saveContactToFile(contactList)
			if err != nil {
				fmt.Println("No se pudo crear el archivo")
				return
			}

			fmt.Println("bander 2")

		case 2:
			// mostrar contactos
			fmt.Println("==========================")
			for index, contacto := range contactList {
				fmt.Printf("%d. Nombre: %s - Email: %s - Telefono: %s \n",
					index, contacto.Name, contacto.Email, contacto.Phone)
			}
			fmt.Println("==========================")
		case 3:
			fmt.Println("Saliendo del programa.. ")
		default:
			break
		}
	}
}
