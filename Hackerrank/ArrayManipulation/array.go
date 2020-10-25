package main

import "fmt"

func arrayManipulation2(n int32, queries [][]int32) int64 {
	//{{1, 3, 2}, {2, 5, 2}, {1, 2, 2}, {4, 6, 2}, {2, 8, 2}},
	helper := make([]int, n+2)
	var max int

	for _, v := range queries {
		if v[0] != v[1] {
			k := int(v[2])
			helper[int(v[0])] += k
			helper[int(v[1])+1] -= k
		}
	}

	fmt.Println(helper)

	for _, v := range helper {
		if v > 0 {
			max += v
		}
	}

	for _, v := range queries {
		if v[0] == v[1] {
			if int(v[2]) > max {
				max = int(v[2])
			}
		}
	}
	return int64(max)

}

///////////////////////////////////////////////

func arrayManipulation(n int32, queries [][]int32) int64 {

	var max int64 = 0
	arr := make([]int64, n+2)
	for _, op := range queries {
		k := int64(op[2])
		arr[op[0]] += k
		arr[op[1]+1] -= k
	}
	var count int64
	for _, v := range arr {
		count += v
		if count > max {
			max = count
		}
	}

	return max

	/*
		No se por que no da
		//{{1, 3, 2}, {2, 5, 2}, {1, 2, 2}, {4, 6, 2}, {2, 8, 2}},
		helper := make([]int32, n+2)
		var max int32

		for _, op := range queries {
			helper[op[0]] += op[2]
			helper[op[1]+1] -= op[2]
		}

		var count int32
		for _, v := range helper {
			count += v
			if count > max {
				max = count
			}
		}

		return int64(max)
	*/
}

/*
//  Solution 1
func arrayManipulation2(n int32, queries [][]int32) int64 {
	helper := make([]int, n+1)
	var max int

	for _, v := range queries {
		origen := int(v[0])
		destino := int(v[1])
		value := int(v[2])

		for j := origen; j <= destino; j++ {
			helper[j] += value
			if helper[j] > max {
				max = helper[j]
			}
		}

	}

	return int64(max)
}
*/
