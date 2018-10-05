package tendencia

import (
	"errors"
)

type Tendencia struct {
	Valores []float32
}

func (t *Tendencia) Promedio() (float32, error) {
	if len(t.Valores) == 0 {
		// La función promedio debería fallar si
		// el array está vacío
		return -1, errors.New("El Array no puede estar vacio")
	}

	var acum float32 = 0
	for _, Valores := range t.Valores {
		acum += Valores
	}
	return acum / float32(len(t.Valores)), nil
}

/*func Moda(el []int) int {
	return 0
}*/

func (t *Tendencia) Mediana() (float32, error) {
	if len(t.Valores) == 0 {
		return -1, errors.New("El Array no puede estar vacio")
	}

	if len(t.Valores)%2 != 0 {
		i := ((len(t.Valores) - 1) / 2)
		return t.Valores[i], nil
	}
	der := len(t.Valores) / 2
	izq := der - 1

	return (t.Valores[der] + t.Valores[izq]) / 2, nil

}
