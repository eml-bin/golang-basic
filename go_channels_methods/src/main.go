package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	// Creación de canal
	c := make(chan string, 2)

	c <- "mensaje_1"
	c <- "mensaje_2"

	fmt.Println(len(c), cap(c))

	// Range y Close
	// Close. cierra un canal para no permitir más entradas de datos
	// Se debe cerrar un channel cuando no hay más datos que entrar..
	// En este escenario se tiene que cerrar para inficarle al runtime de Golang
	// ,que no hay más datos por entrar
	close(c)

	// c <- "mensaje_3"

	// Range para iterar cada uno de los elementos de un canal
	for message := range c {
		fmt.Println(message)
	}

	// Select
	email_1 := make(chan string)
	email_2 := make(chan string)

	go message("mensaje 1", email_1)
	go message("mensaje 2", email_2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email_1:
			fmt.Println("Email Recibido de...", m1)
		case m2 := <-email_2:
			fmt.Println("Email Recibido de...", m2)
		}
	}

}
