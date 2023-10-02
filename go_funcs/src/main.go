package main

import "fmt"

func normalFunction(message string) {
	fmt.Printf("Hello %s", message)
}

// func multipleArgsFunction(a int, b int, c string) {

// Se puede declarar varias variables de un mismo tipo de la siguiente manera
func multipleArgsFunction(a, b int, c string) {
	fmt.Println(a, b, c)
}

// función que retorna un valor, en este caso int
func returnValFunction(num int) int {
	return num * 2
}

// función que retorna dos valores, en este caso ints
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("go!")
	multipleArgsFunction(1, 2, "ABC")

	value := returnValFunction(2)
	println("Value *2: ", value)

	// Ejecutando función que retorna dos valores
	value1, _ := doubleReturn(4)
	println("value1:", value1)
}
