package main

import "fmt"

//https://leetcode.com/problems/minimum-increment-to-make-array-unique/

func minIncrementForUnique(A []int) int {
	// ArrayCount := make(map[int]int, 0)
	ArrayCount := make([]int, 100)
	var max int
	var count int
	for _, value := range A {
		ArrayCount[value]++
		if value > max {
			max = value
		}
	}

	for i, value := range ArrayCount {
		if value > 1 {
			count += value - 1
			ArrayCount[i+1] += value - 1
		}
		fmt.Println(ArrayCount)
	}
	return count
}

/*
Solución 1 : no performa para un input muy grande
func minIncrementForUnique(A []int) int {
	var Increment int
	// fmt.Println("Arrancanding", time.Now())
	for i := 1; i < len(A); i++ {
		// fmt.Println("Entro al primer for con i:", i)
		var IsUnique bool = false
		var ValueAbs int = A[i]
		for IsUnique == false {
			// log.Println("FOR WHILE")
			var marca bool = false
			for j := (i - 1); j >= 0; j-- {
				// fmt.Println("Entro al J for con j:", j)
				if ValueAbs == A[j] {
					// log.Println("Antes del BREAK", i, j)
					ValueAbs++
					Increment++
					marca = true
					break
				}
			}
			if marca == false {
				IsUnique = true
			}
		}
		if ValueAbs != A[i] {
			A[i] = ValueAbs
		}
	}

	return Increment
}
*/
