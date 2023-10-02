// Interfaces en Go. método para compartir un método a otros structs

package main

import "fmt"

type figura2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figura2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	miC := cuadrado{base: 4}
	miR := rectangulo{base: 4, altura: 3}

	// Implementación de <interface> `figura2D`
	calcular(miC)
	calcular(miR)

	// Lista de Interfaces
	miInterface := []interface{}{"ABC", 123, 0.5}
	fmt.Println(miInterface...)

}
