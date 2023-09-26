package main

import "fmt"

func main() {
	// Área cuadrado

	const baseCuadrado = 10

	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Área cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	// División
	result = y / x
	fmt.Println("División: ", result)

	// Módulo
	result = y % x
	fmt.Println("División: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Números complejos
	c := 10 + 8i
	fmt.Println("Complex: ", c)
}
