package arrays

import (
	"errors"
)

var array []int{1 ,3, 6, 8, 9}

func Get(i int) (int, error) {

	if i < 0 || i > len(array)-1 {
		return -1, errors.New("El indice es invalido")
	}
	return array[i], nil
}
