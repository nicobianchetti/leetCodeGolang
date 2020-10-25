package main

import (
	"testing"
)

func TestUnique(t *testing.T) {

	casos := []struct {
		nombre  string
		arrayIn []int
		result  bool
	}{
		{
			nombre:  "Ejemplo 1",
			arrayIn: []int{1, 2, 2, 1, 1, 3},
			result:  true,
		},
		{
			nombre:  "Ejemplo 2",
			arrayIn: []int{1, 2},
			result:  false,
		},
		{
			nombre:  "Ejemplo 3",
			arrayIn: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			result:  true,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			// resultF := twoSum(test.arrayIn, test.tarjetIn)
			Result := uniqueOccurrences(test.arrayIn)

			if Result != test.result {
				t.Errorf("Retorna valores incorrectos")
			}
		})
	}

}
