package main

import "testing"

func TestDeleteColum(t *testing.T) {
	casos := []struct {
		nombre string
		A      []string
		result int
	}{
		{
			nombre: "Case 1:`[cba,daf,ghi]`",
			A:      []string{"cba", "daf", "ghi"},
			result: 1,
		},

		{
			nombre: "Case 2:`[cba,daf,ghi]`",
			A:      []string{"a", "b"},
			result: 0,
		},

		{
			nombre: "Case 3:`[cba,daf,ghi]`",
			A:      []string{"zyx", "wvu", "tsr"},
			result: 3,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			D := minDeletionSize(test.A)

			if D != test.result {
				t.Error("Retorna", D, ",y el valor correcto es", test.result)
			}
		})
	}

}
