package main

import (
    "github.com/emikohmann/ubp-lab3/ejemplos/estructuras/alumnos"
    "fmt"
)

func main() {

    alumno1 := alumnos.Alumno{
        Nombre: "Juancito",
        Legajo: 1234567,
        Promedio: 7.67,
    }

    fmt.Println("Nombre: ", alumno1.Nombre)

    // modificar(alumno1)

    alumno1 = alumno1.Modificar()

    fmt.Println("Nombre nuevo: ", alumno1.Nombre)
}

/*
func modificar(al *alumnos.Alumno) {

    al.Nombre = "Lucas"
}
*/