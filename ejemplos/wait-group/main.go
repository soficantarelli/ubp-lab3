package main

import (
	"fmt"
	"sync"
)

func main() {

	//creo un wait group
	var wg sync.WaitGroup

	//informmo al wait group que voy a ejecutar 2 go routines
	wg.Add(2)

	go mostrarMensaje("HOLA", &wg)
	go mostrarMensaje("CHAU", &wg)

	//hacemos que espere la finalizacion de las go routines
	wg.Wait()

}

func mostrarMensaje(msj string, wg *sync.WaitGroup) {

	fmt.Println(msj)

	//informo el wait group que la go routine termino
	wg.Done()
}
