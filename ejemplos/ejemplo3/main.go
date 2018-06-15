//mediana y moda

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []float64{9, 7.5, 10, 2, 3, 1, 5, 5}

	sort.Float64s(a)
	fmt.Println("Despues de ordenar: ", a)
	fmt.Println("Mediana: ", mediana(a))
	fmt.Println("Moda: ", moda(a))
}

func mediana(vs []float64) float64 {
	var aux = len(vs) / 2
	if len(vs)%2 == 0 {
		return float64((vs[aux-1] + vs[aux]) / 2)
	}
	return vs[aux]
}

func moda(vs []float64) float64 {
	maxrepeticiones := 0
	moda := 0.0
	n := len(vs)
	for i := 0; i < n; i++ {
		var rep = 0
		for j := 0; j < n; j++ {
			if vs[i] == vs[j] {
				rep++
			}
			if rep > maxrepeticiones {
				moda = vs[j]
				maxrepeticiones = rep
			}
		}
	}

	return moda
}
