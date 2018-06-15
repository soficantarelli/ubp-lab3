//promedio

package main

import (
	"fmt"
	"errors"
)

func main() {
	//var valores []int
	valores := []float32{
		3, 4, 5.2, 9,
	}
	fmt.Println(promedio(valores))
}

func promedio(vs []float32) (float32,error) {

	if len(vs) == 0{
		return -1, errors.New()
	}

	var acum float32 = 0

	for _, v := range vs { // i = indice pero como no se usa se pone _, v = valor
		acum += v
	}
	return acum / float32(len(vs))
}

/*func promedio(vs []float32) float32 {

	var acum float32 = 0
	cont := 0

	tam := len(vs)

	for i := 0; i < tam; i++ {

		acum += vs[i]
		cont++
	}
	return acum / float32(cont)
}*/
