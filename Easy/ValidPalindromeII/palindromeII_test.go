package main

import (
	"testing"
)

func TestAdvantage(t *testing.T) {
	casos := []struct {
		nombre string
		in     string
		result bool
	}{
		{
			nombre: "Case 1:",
			in:     "aba",
			result: true,
		},

		{
			nombre: "Case 2:",
			in:     "abca",
			result: true,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			result := validPalindrome(test.in)

			if result == test.result {
				t.Error("Retorna", result, ",y el valor correcto es", test.result)
			}
		})
	}

}
