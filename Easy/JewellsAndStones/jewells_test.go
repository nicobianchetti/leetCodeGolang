package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	casos := []struct {
		nombre string
		jIn    string
		sIn    string
		result int
	}{
		{
			nombre: "Obtiene la cantida de joyas dentro de las piedras que ya tengo",
			jIn:    "aA",
			sIn:    "aAAbbbb",
			result: 3,
		},

		{
			nombre: "Pruebo distinción Mayúsculas y minúsculas ",
			jIn:    "z",
			sIn:    "ZZ",
			result: 0,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			ResultJewells := numJewelsInStones(test.jIn, test.sIn)

			if ResultJewells != test.result {
				t.Error("Retorna", ResultJewells, ",y el valor correcto es", test.result)
			}
		})
	}

}
