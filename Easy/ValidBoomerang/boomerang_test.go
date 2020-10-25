package main

import (
	"testing"
)

func TestAdvantage(t *testing.T) {
	casos := []struct {
		nombre string
		in     [][]int
		result bool
	}{
		{
			nombre: "Case 1:",
			in:     [][]int{{1, 1}, {2, 3}, {3, 2}},
			result: true,
		},

		{
			nombre: "Case 2:",
			in:     [][]int{{1, 1}, {2, 2}, {3, 3}},
			result: false,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			result := isBoomerang(test.in)

			if result == test.result {
				t.Error("Retorna", result, ",y el valor correcto es", test.result)
			}
		})
	}

}
