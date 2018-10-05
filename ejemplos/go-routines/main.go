package main

import "fmt"

func main() {

	//fmt.Println("Hola mundo")

	/*

		go fmt.Println("Hola mundo")

		for true {}

	*/

	for i := 0; i < 100; i++ {
		go mostrarMensaje(i)
	}

	//go mostrarMensaje(56)

}

func mostrarMensaje(i int) {

	msg := fmt.Sprintf("Hola mundo %d", i)
	fmt.Println(msg)
}
