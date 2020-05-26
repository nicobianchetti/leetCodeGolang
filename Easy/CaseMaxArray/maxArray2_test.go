package main

import "testing"

func TestCaseMaxArray2(t *testing.T) {
	casos := []struct {
		name   string
		In     []int
		result int
	}{
		{
			name:   "Case 1",
			In:     []int{1, 2, 4, 6, 4, 12, 3, 234, 231, 4, 5, 1, 3, 5, 6, 7, 23, 234, 99999999999, 245, 6456, 7457, 45645, 44444, 6, 5, 4, 7, 9, 4, 1, 2, 3123123, 123123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123},
			result: 99999999999,
		},
	}

	for _, test := range casos {
		t.Run(test.name, func(t *testing.T) {
			Result := MaxArrayCase2(test.In)

			if Result != test.result {
				t.Error("Retorna", Result, ",y el valor correcto es", test.result)
			}
		})
	}
}
