package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("es 1")
	} else {
		fmt.Println("no es 1")
	}

	// AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Sí valor1=1 Y(AND) valor2=2")
	}

	// AND
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Si valor1=1 Y(AND) valor2=3")
	}

	// OR
	if valor1 == 1 || valor2 == 3 {
		fmt.Println("Si valor1=1 O(OR) valor2=3")
	}

	// CAST string to int
	// value, err := strconv.Atoi("53") // strconv. Retorna dos valores (value. valor convertido | err. si existe un error lo guarda aquí )
	value, err := strconv.Atoi("ad3") // strconv con error

	// Si no existe un valor, es decir, si el error es nil
	if err != nil {
		log.Fatal(err)
	}

	println("CAST Value:", value)

}
