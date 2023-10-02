package main

import (
	"fmt"
	"sync"
	"time"
)

// Go Concurrencia.
//

func say(text string, wg *sync.WaitGroup) {

	// ültima acción del bloque, el fin del waitgroup
	defer wg.Done()

	fmt.Println(text)
}

func main() {

	// `sync.WaitGroup`. Acumula un grupo de go-routines, para liberar poco a poco
	var wg sync.WaitGroup

	fmt.Println("Hello")

	// Se agrega una go-routine
	wg.Add(1)

	go say("Wola", &wg)

	// Indicarle al Wait group que espere a que todas la go-routine finalicen
	wg.Wait()

	// Función anónima
	go func(text string) {
		fmt.Println(text)
	}("Adeus")

	// En este escenario se puede utilizar time.Sleep para permitir ejecutar las go-routines
	// pero no es la mejor implementación
	time.Sleep(time.Second * 1)

}
