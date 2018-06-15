package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	n := 5
	print(" Bosque creado \n")
	m := crearBosque(n)

	print(" Cargar bosque \n")
	cargarBosque(m)
	mostrarBosque(m)

}

func crearBosque(n int) [][]int {

	matriz := make([][]int, n)
	for i := range matriz {
		matriz[i] = make([]int, n)
		for j := range matriz[i] {
			matriz[i][j] = 0
		}
	}
	return matriz
}

func cargarBosque(matriz [][]int) {

	posicion1 := 0
	posicion2 := 0

	rand.Seed(time.Now().Unix())

	for i := 0; i <= (len(matriz)*len(matriz))/2; i++ {
		posicion1 = rand.Intn(len(matriz))
		posicion2 = rand.Intn(len(matriz))
		matriz[posicion1][posicion2] = rand.Intn(2) + 1
	}
}

func mostrarBosque(matriz [][]int) {

	fmt.Println(matriz)
}
