package main

import "fmt"

func main() {
	// Declaración de constantes
	// const - nunca cambia de valor en el tiempo

	const pi float64 = 3.1416 // variable tipada
	const pi2 = 3.14

	fmt.Println("pi", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12 // utilizando `:` indica que la variable no ha sido declarada y Go se encarga de crearla
	var altura int = 14
	var area int = 12 * 14

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Área de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Área de Cuadrado: ", areaCuadrado)
}
