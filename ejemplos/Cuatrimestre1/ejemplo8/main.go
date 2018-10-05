//promedio, errores

package main

import (
	"errors"
	"fmt"
)

func main() {
	//var valores []int
	valores, err := promedio([]float32{})
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println(promedio)
}

func promedio(vs []float32) (float32, error) {

	if len(vs) == 0 {
		return -1, errors.New("El array no puede estar vacio")
	}

	var acum float32 = 0

	for _, v := range vs { // i = indice pero como no se usa se pone _, v = valor
		acum += v
	}
	return acum / float32(len(vs)), nil
}
