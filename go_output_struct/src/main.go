package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

// Añadir nuevo método a <struct> `pc` String() que reemplaza el default de un Println
// x. representa el struct
// String(). nombre de la función añadida
// string. tipo de dato que retorna
func (x pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es marca %s", x.ram, x.disk, x.brand)
}

func main() {
	miPC := pc{ram: 16, brand: "msi", disk: 100}

	fmt.Println(miPC)
}
