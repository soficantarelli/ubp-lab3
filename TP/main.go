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

	print("\n")

	ejecutarSimulacion(m)
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

	for _, matriz := range matriz {
		fmt.Println(matriz)
	}

}

func ejecutarSimulacion(matriz [][]int) {

	vueltasTotales

	ultimaMatriz := matriz

	for true {
		actualizarUltimoBosque(ultimaMatriz, matriz)
		ejecutarPaso(matriz, vueltasTotales)
		if terminoCiclo(ultimaMatriz, matriz) {
			fmt.Println("\nVueltas totales:", vueltasTotales-1)
			return
		}
		mostrarBosque(matriz)
	}
}

func actualizarUltimoBosque(ultimaMatriz [][]int, matriz [][]int) {
	for i, f := range matriz {
		for j := range f {
			ultimaMatriz[i][j] = matriz[i][j]
		}
	}
}

func terminoCiclo(ultimaMatriz [][]int, matriz [][]int) bool {
	for i, f := range matriz {
		for j := range f {
			if ultimaMatriz[i][j] != matriz[i][j] {
				return false
			}
		}
	}
	return true
}

func ejecutarPaso(matriz [][]int, vueltasTotales int) {

	vueltasTotales++

	for i, f := range matriz {
		for j, c := range f {

			switch c {
			case 0:
				break

			case 1:
				if hayZorro(matriz, i, j) {
					matriz[i][j] = 0
				} else {
					crearNuevo(matriz, 1)
				}
				break

			case 2:
				if hayHierba(matriz, i, j) {
					matriz[i][j] = 0
				} else {
					crearNuevo(matriz, 2)
				}
				break
			}
		}
	}
}

func hayZorro(m [][]int, i int, j int) bool {

	si, sc, sd := 0, 0, 0
	ii, ic, id := 0, 0, 0
	izq, der := 0, 0

	if j > 0 {

		if i > 0 {
			si = m[i-1][j-1]
		}

		sc = m[i][j-1]

		if i < len(m)-1 {
			sd = m[i+1][j-1]
		}
	}

	if i > 0 {
		izq = m[i-1][j]
	}

	if i < len(m)-1 {
		der = m[i+1][j]
	}

	if j < len(m)-1 {
		if i > 0 {
			ii = m[i-1][j+1]
		}

		ic = m[i][j+1]
		if i < len(m)-1 {
			id = m[i+1][j+1]
		}
	}

	if si == 2 || sc == 2 || sd == 2 || izq == 2 || der == 2 || ii == 2 || ic == 2 || id == 2 {
		return true
	}
	return false

}

func hayHierba(m [][]int, i int, j int) bool {

	si, sc, sd := 0, 0, 0
	ii, ic, id := 0, 0, 0
	izq, der := 0, 0

	if j > 0 {

		if i > 0 {
			si = m[i-1][j-1]
		}

		sc = m[i][j-1]

		if i < len(m)-1 {
			sd = m[i+1][j-1]
		}
	}

	if i > 0 {
		izq = m[i-1][j]
	}

	if i < len(m)-1 {
		der = m[i+1][j]
	}

	if j < len(m)-1 {
		if i > 0 {
			ii = m[i-1][j+1]
		}

		ic = m[i][j+1]
		if i < len(m)-1 {
			id = m[i+1][j+1]
		}
	}

	if si == 0 && sc == 0 && sd == 0 && izq == 0 && der == 0 && ii == 0 && ic == 0 && id == 0 {
		return true
	}
	return false

}

func crearNuevo(m [][]int, elemento int) {
	n := len(m)
	m[rand.Intn(n)][rand.Intn(n)] = elemento
}
