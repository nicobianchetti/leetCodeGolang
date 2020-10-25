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
			in:     "A man, a plan, a canal: Panama",
			result: true,
		},

		{
			nombre: "Case 2:",
			in:     "race a car",
			result: false,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			result := isPalindrome(test.in)

			if result == test.result {
				t.Error("Retorna", result, ",y el valor correcto es", test.result)
			}
		})
	}

}
