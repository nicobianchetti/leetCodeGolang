package main

import (
	"testing"
)

func TestAdvantage(t *testing.T) {
	casos := []struct {
		nombre string
		s      string
		t      string
		result bool
	}{
		{
			nombre: "Case 1:",
			s:      "anagram",
			t:      "nagaram",
			result: true,
		},

		{
			nombre: "Case 2:",
			s:      "rat",
			t:      "car",
			result: false,
		},

		{
			nombre: "Case 3:",
			s:      "a",
			t:      "b",
			result: false,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			result := isAnagram(test.s, test.t)

			if result != test.result {
				t.Error("Retorna", result, ",y el valor correcto es", test.result)
			}
		})
	}

}
