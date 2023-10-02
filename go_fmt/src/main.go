package main

import "fmt"

func main() {
	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "Go!"

	// Println -> print con salto de línea al final
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf -> print con un formato
	nombre := "Devs"
	cantidad := 1000
	fmt.Printf("%s tiene más de %d habitantes, %v\n", nombre, cantidad, "any-val")
	// Buena práctica es saber el tipo de dato pero existe %v para esos casos donde no

	// Sprintf -> generar un string sin imprimir en consola
	message := fmt.Sprintf("%s tiene más de %d habitantes, %v", nombre, cantidad, "")
	fmt.Println(message)

	// Conocer tipo de dato. ejemplos
	fmt.Printf("message(type): %T\n", message)
	fmt.Printf("cantidad(type): %T\n", cantidad)
}
