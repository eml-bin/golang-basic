package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {

	text = strings.ToUpper(text)
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("SI es palindromo")
	} else {
		fmt.Println("NO es palindromo")
	}
}

func main() {
	slice := []string{"hola", "mundo", "you"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	isPalindromo("casa")
	isPalindromo("ana")
	isPalindromo("ama")
	isPalindromo("amA")
}
