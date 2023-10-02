package main

import "fmt"

func main() {
	// Defer. keyword para ejecutar una instrucción antes de morir
	// Se puede utilizar por ejemplo cerrar una conexión
	// Buena práctica utilizar un solo defer por bloque de código
	defer fmt.Println("hola")
	fmt.Println("mundo")

	x := 1
	fmt.Println(x)

	// Continue & break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// continue
		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		if i == 6 {
			fmt.Println("salir con break")
			break
		}
	}

	for {
		fmt.Println("Ciclo \"infinito\"")
		break
	}
}
