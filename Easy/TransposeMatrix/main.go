package main

/*
https://leetcode.com/problems/transpose-matrix/

[1,2,3]           [1,4,7]
[4,5,6]    ---->  [2,5,8]
[7,8,9]           [3,6,9]
*/

func transpose(A [][]int) [][]int {

	lenI := len(A)
	lenJ := len(A[0])

	//Solution 1
	// Trasp := [][]int{}
	// for j := 0; j < lenJ; j++ {
	// 	B := []int{}
	// 	for i := 0; i < lenI; i++ {
	// 		B = append(B, A[i][j])
	// 	}
	// 	Trasp = append(Trasp, B)
	// }

	//Solution 2
	Trasp := make([][]int, lenJ)

	for i := range Trasp {
		Trasp[i] = make([]int, lenI)
		for j := range Trasp[i] {
			Trasp[i][j] = A[j][i]
		}
	}

	return Trasp
}
