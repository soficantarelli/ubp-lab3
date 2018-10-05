//ordenar

package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []float64{9, 7, 8, 5, 3, 2, 4, 1, 6}

	fmt.Println("Antes de ordenar %v", a)
	sort.Float64s(a)
	fmt.Println("Despues de ordenar: ", a)

}
