package main

import "testing"

func TestDeleteColum(t *testing.T) {
	casos := []struct {
		nombre string
		A      []int
		result int
	}{
		// {
		// 	nombre: "Case 1:[1,2,2]",
		// 	A:      []int{1, 2, 2},
		// 	result: 1,
		// },

		// {
		// 	nombre: "Case 2:[3,2,1,2,1,7]",
		// 	A:      []int{3, 2, 1, 2, 1, 7},
		// 	result: 6,
		// },

		// {
		// 	nombre: "Case 3:[3,2,1,2,1,7]",
		// 	A:      []int{1, 2, 3, 2, 3, 4, 3, 1},
		// 	result: 17,
		// },

		{
			nombre: "Case 4:[3,2,1,2,1,7]",
			A:      []int{3, 2, 1, 2, 1, 7, 3, 3, 4, 4},
			result: 25,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			D := minIncrementForUnique(test.A)

			if D != test.result {
				t.Error("Retorna", D, ",y el valor correcto es", test.result)
			}
		})
	}

}
