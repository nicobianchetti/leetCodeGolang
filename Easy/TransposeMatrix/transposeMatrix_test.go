package main

import (
	"reflect"
	"testing"
)

func TestProjectionArea(t *testing.T) {
	casos := []struct {
		name   string
		matrix [][]int
		result [][]int
	}{
		{
			name:   "Case 1:[[1,2,3],[4,5,6],[7,8,9]]",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			result: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},

		{
			name:   "Case 2:[[1,2,3],[4,5,6]]",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}},
			result: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
	}

	for _, test := range casos {
		t.Run(test.name, func(t *testing.T) {
			Result := transpose(test.matrix)

			if !reflect.DeepEqual(Result, test.result) {
				t.Error("Retorna", Result, ",y el valor correcto es", test.result)
			}
		})
	}
}
