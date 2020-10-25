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
			nombre:  "Case 1:[1,2,2]",
			A:       []int32{2, 1, 5, 3, 4},
			result:  3,
			withErr: false,
		},

		{
			nombre:  "Case 2:",
			A:       []int32{2, 5, 1, 3, 4},
			result:  0,
			withErr: true,
		},

		{
			nombre:  "Case 3:",
			A:       []int32{5, 1, 2, 3, 7, 8, 6, 4},
			result:  0,
			withErr: true,
		},

		{
			nombre:  "Case 4:",
			A:       []int32{1, 2, 5, 3, 7, 8, 6, 4},
			result:  7,
			withErr: false,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			result, error := minimumBribes(test.A)

			switch {
			case result != test.result:
				t.Errorf("Se esperaba un retorno de %d, y se retornó %d", test.result, result)
			// case test.withErr && error == nil:
			case test.withErr != error:
				t.Errorf("Se esperaba un caos, pero se recibió vacío")
			}
		})
	}

}
