package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	casos := []struct {
		nombre string
		num    int
		result int
	}{
		{
			nombre: "Obtiene la cantida de pasos necesarios para restar el nro hasta llegar a 0",
			num:    123,
			result: 12,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			// resultF := twoSum(test.arrayIn, test.tarjetIn)
			ResultStep := numberOfSteps(test.num)

			if ResultStep == test.num {
				t.Error("Retorna", ResultStep, ",y el valor correcto es", test.num)
			}
		})
	}

}
