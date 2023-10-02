package main

import "fmt"

func main() {
	// modulo := 4 % 2

	// Switch case
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Switch case sin condición
	value := 50
	switch {
	case value > 100:
		fmt.Println("mayor a 100")
	case value < 0:
		fmt.Println("es negativo")
	default:
		// cuando es menor de 100 pero mayor a 0
		fmt.Println("Sin condición")
	}
}
