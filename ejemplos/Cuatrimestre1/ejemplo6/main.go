package main

import (
	"fmt"

	"github.com/soficantarelli/ubp-lab3/ejemplos/ejemplo5/alumnos"
)

func main() {

	alumno1 := alumnos.Alumno{
		Nombre:   "Juan",
		Legajo:   5892,
		Promedio: 7.68,
	}
	fmt.Println("Nombre: ", alumno1.Nombre)
	modificar(&alumno1)
	fmt.Println("Nombre nuevo: ", alumno1.Nombre)
}

func (al *Alumno) modificar() {
	al.Nombre = "Lucas"
}
