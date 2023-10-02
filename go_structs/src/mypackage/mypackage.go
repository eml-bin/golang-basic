// Custom Paquete
package mypackage

import "fmt"

// CarPublic `struct` Car con acceso público, usando máyuscula al principio es público
type CarPublic struct {
	Brand string
	Year  int
}

// carPrivate `struct` Car con acceso privado, usando minúsculas al principio es privado
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage. Func Pública
func PrintMessage(text string) {
	fmt.Println("Hello PK", text)
}

// PrintMessage. Func Privada
func printMessage1(text string) {
	fmt.Println("Hello PK", text)
}
