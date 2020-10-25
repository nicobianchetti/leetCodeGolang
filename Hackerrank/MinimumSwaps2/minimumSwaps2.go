package main

/*
https://www.hackerrank.com/challenges/minimum-swaps-2/problem

Sample Input 0

4
4 3 1 2
Sample Output 0

3

Sample Input 1

5
2 3 4 1 5
Sample Output 1

3

*/

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	length := len(arr)
	helper := make([]int, length+1)
	var isOrder bool
	var moves int32
	for i, v := range arr {
		helper[v] = i
	}

	for i := 1; i < length; i++ {
		isOrder = false
		for isOrder == false {
			// fmt.Println("proceso", i)
			if helper[i] != i-1 {
				moves++
				arrOrigen := arr[helper[i]]
				arrDestino := arr[i-1]
				helperOrigen := helper[i]
				helperDestino := helper[arrDestino]
				arr[i-1], arr[helper[i]] = arrOrigen, arrDestino
				helper[arrDestino], helper[i] = helperOrigen, helperDestino
				isOrder = true
			} else {
				isOrder = true
			}
		}
	}

	return moves
}
