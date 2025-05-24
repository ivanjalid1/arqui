package main

import "fmt"

func main() {
	// Valores predeterminados

	var (
		defaultInt     int
		defaultUint    uint
		defaultFloat32 float32
		defaultBool    bool
		defaultString  string
	)

	fmt.Println(defaultInt, defaultUint, defaultFloat32, defaultBool, defaultString)

	var num1 float32 = 2.52
	var num2 int16 = 14

	fmt.Println(num2 + int16(num1))
}
