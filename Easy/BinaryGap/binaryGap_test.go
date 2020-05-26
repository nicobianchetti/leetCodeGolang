package main

import "testing"

func TestBinaryGap(t *testing.T) {
	casos := []struct {
		name   string
		In     int
		result int
	}{
		{
			name:   "Case 1",
			In:     22,
			result: 2,
		},

		// {
		// 	name:   "Case 2",
		// 	In:     5,
		// 	result: 2,
		// },

		// {
		// 	name:   "Case 3",
		// 	In:     6,
		// 	result: 1,
		// },

		// {
		// 	name:   "Case 4",
		// 	In:     8,
		// 	result: 0,
		// },
	}

	for _, test := range casos {
		t.Run(test.name, func(t *testing.T) {
			Result := binaryGap(test.In)

			if Result != test.result {
				t.Error("Retorna", Result, ",y el valor correcto es", test.result)
			}
		})
	}
}
