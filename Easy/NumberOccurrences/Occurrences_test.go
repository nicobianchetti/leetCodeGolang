package main

import (
	"testing"
)

func TestOccurrences(t *testing.T) {

	casos := []struct {
		nombre  string
		arrayIn []int
		result  bool
	}{
		{
			nombre:  "Case 1: verifica si el nro de ocurrencias de los elementos del array es único",
			arrayIn: []int{1, 2, 2, 1, 1, 3},
			result:  true,
		},

		{
			nombre:  "Case 2: verifica si el nro de ocurrencias de los elementos del array es único",
			arrayIn: []int{1, 2},
			result:  false,
		},

		{
			nombre:  "Case 3: verifica si el nro de ocurrencias de los elementos del array es único",
			arrayIn: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			result:  true,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			ResultOccurrences := uniqueOccurrences(test.arrayIn)

			if ResultOccurrences != test.result {
				t.Error("Retorna", ResultOccurrences, ",y el valor correcto es", test.result)
			}
		})
	}
}
