package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (miPC pc) ping() {
	fmt.Println(miPC.brand, "Pong")
}

func (miPC *pc) duplicateRAM() {
	miPC.ram = miPC.ram * 2
}

func main() {
	a := 50

	// &a. sería el puntero (dirección de memoria) de donde vive a
	b := &a

	fmt.Println(b)

	// *b. accede al valor del puntero
	fmt.Println(*b)

	// Reasignar valor del puntero
	*b = 100
	fmt.Println(a)

	// otros usos de punteros

	miPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(miPC)

	miPC.ping()

	fmt.Println(miPC)
	miPC.duplicateRAM()

	fmt.Println(miPC)
	miPC.duplicateRAM()

	fmt.Println(miPC)

}
