package main

import "fmt"

func main() {
	// ARRAY
	var array [4]int
	fmt.Println(array)

	array[0] = 1
	array[1] = 2
	// len. cuenta el elementos actuales de un array
	// cap. cuenta la capacidad de un array
	fmt.Println(array, len(array), cap(array))

	// SLICE

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el SLICE

	// Slicing
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append, nuevo elemento
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append, nueva list
	newSlice := []int{8, 9, 10}
	// los `...` indican que golang agregara los items de la otra lista
	slice = append(slice, newSlice...)

}
