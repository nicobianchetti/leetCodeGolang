package main

import (
	"reflect"
	"testing"
)

func TestAdvantage(t *testing.T) {
	casos := []struct {
		nombre string
		A      []int
		B      []int
		result []int
	}{
		// {
		// 	nombre: "Case 1:[1,2,2]",
		// 	A:      []int{2, 7, 11, 15},
		// 	B:      []int{1, 10, 4, 11},
		// 	result: []int{2, 11, 7, 15},
		// },

		// {
		// 	nombre: "Case 2:",
		// 	A:      []int{12, 24, 8, 32},
		// 	B:      []int{13, 25, 32, 11},
		// 	result: []int{24, 32, 8, 12},
		// },

		{
			nombre: "Case 3:",
			A:      []int{2, 0, 4, 1, 2},
			B:      []int{1, 3, 0, 0, 2},
			result: []int{2, 0, 2, 1, 4},
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			result := advantageCount(test.A, test.B)

			if !reflect.DeepEqual(result, test.result) {
				t.Error("Retorna", result, ",y el valor correcto es", test.result)
			}
		})
	}

}
