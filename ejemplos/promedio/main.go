package main

import "fmt"

func main() {
    valores := []float32{
        3.13,4.02,5.5,6.32,7.43,8.24,
    }

    fmt.Println("Promedio: ", promedio(valores))
}

/*
func promedio(vs []float32) float32 {
    var acum float32 = 0
    cont := 0

    tam := len(vs)

    for i := 0; i < tam; i++ {
        acum = acum + vs[i]
        cont = cont + 1
    }

    return acum / float32(cont)
}
*/

func promedio(vs []float32) float32 {
    var acum float32 = 0
    for _, v := range vs {
        acum += v
    }
    return acum / float32(len(vs))
}