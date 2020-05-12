package main

import "testing"

func TestProjectionArea(t *testing.T) {
	casos := []struct {
		name   string
		girdIn [][]int
		result int
	}{
		{
			name:   "Case 1:[[2]]",
			girdIn: [][]int{{2}},
			result: 5,
		},

		{
			name:   "Case 2:[[1,2],[3,4]]",
			girdIn: [][]int{{1, 2}, {3, 4}},
			result: 17,
		},

		{
			name:   "Case 3:[[1,0],[0,2]]",
			girdIn: [][]int{{1, 0}, {0, 2}},
			result: 8,
		},

		{
			name:   "Case 4:[1,1,1],[1,0,1],[1,1,1]",
			girdIn: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			result: 14,
		},

		{
			name:   "Case 5:[[2,2,2],[2,1,2],[2,2,2]]",
			girdIn: [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}},
			result: 21,
		},

		{
			name:   "Case 6:[[2,3],[2,4]]",
			girdIn: [][]int{{2, 3}, {2, 4}},
			result: 17,
		},
	}

	for _, test := range casos {
		t.Run(test.name, func(t *testing.T) {
			ResultProjections := projectionArea(test.girdIn)

			if ResultProjections != test.result {
				t.Error("Retorna", ResultProjections, ",y el valor correcto es", test.result)
			}
		})
	}
}
