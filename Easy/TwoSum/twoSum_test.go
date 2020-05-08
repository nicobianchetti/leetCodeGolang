package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	casos := []struct {
		nombre   string
		arrayIn  []int
		tarjetIn int
		result   []int
	}{
		{
			nombre:   "Obtiene el par de índices correctos",
			arrayIn:  []int{2, 7, 11, 15},
			tarjetIn: 9,
			result:   []int{0, 1},
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			// resultF := twoSum(test.arrayIn, test.tarjetIn)
			ResultArray := twoSum(test.arrayIn, test.tarjetIn)

			if !reflect.DeepEqual(ResultArray, test.result) {
				t.Errorf("Retorna valores incorrectos")
			}
		})
	}

}
