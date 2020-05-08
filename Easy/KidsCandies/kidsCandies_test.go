package main

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	casos := []struct {
		nombre       string
		candies      []int
		extraCandies int
		result       []bool
	}{
		{
			nombre:       "Case 1:Obtiene kids que puede llegar al maximo de candies según extra candies",
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			result:       []bool{true, false, true},
		},

		{
			nombre:       "Case 2:Obtiene kids que puede llegar al maximo de candies según extra candies",
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			result:       []bool{true, false, false, false, false},
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			resultCandies := kidsWithCandies(test.candies, test.extraCandies)

			if !reflect.DeepEqual(resultCandies, test.result) {
				t.Error("Retorna", resultCandies, ",y el valor correcto es", test.result)
			}
		})
	}

}
