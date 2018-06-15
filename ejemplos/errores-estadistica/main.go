package main

import (
	"fmt"

	"github.com/soficantarelli/ubp-lab3/ejemplos/errores-estadistica/tendencia"
)

func main() {

	t := &tendencia.Tendencia{
		Valores: []float32{},
	}

	prom, err := t.Promedio()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("Promedio: ", prom)

	/*moda, err := t.Moda()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
	    fmt.Println("Moda: ", moda)*/

	mediana, err := t.Mediana()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("Mediana: ", mediana)

}
