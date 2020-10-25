package main

import (
	"testing"
)

func TestAdvantage(t *testing.T) {
	casos := []struct {
		nombre string
		prices []int32
		k      int32
		result int32
	}{
		{
			nombre: "Case 1:",
			prices: []int32{1, 2, 3, 4},
			k:      7,
			result: 3,
		},

		{
			nombre: "Case 2:",
			prices: []int32{1, 12, 5, 111, 200, 1000, 10},
			k:      50,
			result: 4,
		},

		{
			nombre: "Case 3:",
			prices: []int32{49, 111, 200, 1000},
			k:      50,
			result: 1,
		},

		{
			nombre: "Case 4:",
			prices: []int32{111, 200, 1000},
			k:      50,
			result: 0,
		},

		{
			nombre: "Case 5:",
			prices: []int32{3, 7, 2, 9, 4},
			k:      15,
			result: 3,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			result := maximumToys(test.prices, test.k)

			if result != test.result {
				t.Error("Retorna", result, ",y el valor correcto es", test.result)
			}
		})
	}

}
