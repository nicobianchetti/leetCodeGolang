package main

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

/*
la funcion recibe como parametros un arreglo de enteros y un target que representa la suma
de dos numeros que integran el arreglo.Son siempre dos numeros distintos. Devuelve un arreglo con
los indices donde se encuentran esos dos sumandos
*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}
