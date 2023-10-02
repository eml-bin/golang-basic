package main

import (
	"fmt"
	pk "go_structs/src/mypackage"
)

// struct. La forma de crear clases en Golang
type car struct {
	brand string
	year  int
}

func main() {

	// Instanciar struct
	myCar := car{brand: "ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera de instanciar
	var otherCar car
	otherCar.brand = "Ferrari"

	fmt.Println(otherCar)

	// USando custom packages

	var customCar pk.CarPublic
	customCar.Brand = "Honda"
	customCar.Year = 2000
	fmt.Println(customCar)

	// var customPrivateCar pk.carPrivate
	// fmt.Println(customPrivateCar)
	// customPrivateCar.brand = "t"

	pk.PrintMessage("Bye")
	// pk.printMessage1("Bye")

}
