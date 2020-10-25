package main

/*
https://www.hackerrank.com/challenges/mark-and-toys/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting

*/

import (
	"sort"
)

func maximumToys(prices []int32, k int32) int32 {

	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })

	var total int32
	isExist := false

	for i, v := range prices {
		if v <= k {
			isExist = true
		}
		total += v
		if total > k {
			return int32(i)
		}
	}

	if isExist == false {
		return 0
	}
	return int32(len(prices))

}
