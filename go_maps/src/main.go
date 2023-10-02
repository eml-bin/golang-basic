package main

import "fmt"

func main() {
	// Maps

	// se utilizan make, que sirve para crear diccionarios y otro tipos de objetos
	m := make(map[string]int)

	m["jose"] = 15
	m["maria"] = 18

	// imprimir map
	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i)
		fmt.Println(v)
	}

	// Acceso a un valor
	// ok es un valor que indica si existe o no existe en el map (diccionario)
	value, ok := m["jose"]
	fmt.Println(value, ok)
}
