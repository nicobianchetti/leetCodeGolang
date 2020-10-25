package main

import (
	"testing"
)

func TestAdvantage(t *testing.T) {
	casos := []struct {
		nombre  string
		A       []int32
		result  int
		withErr bool
	}{
		{
			nombre: "Case 1:",
			A:      []int32{4, 3, 1, 2},
			result: 3,
		},

		{
			nombre: "Case 2:",
			A:      []int32{2, 3, 4, 1, 5},
			result: 3,
		},

		{
			nombre: "Case 2:",
			A:      []int32{1, 3, 5, 2, 4, 6, 7},
			result: 3,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			result := minimumSwaps(test.A)

			if int(result) != test.result {
				t.Errorf("Se esperaba un retorno de %d, y se retornó %d", test.result, result)
			}
		})
	}

}
