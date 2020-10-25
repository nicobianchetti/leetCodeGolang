package main

import (
	"sort"
)

/*

https://leetcode.com/problems/advantage-shuffle/

Given two arrays A and B of equal size, the advantage of A with respect to B is the number of indices i for which A[i] > B[i].

Return any permutation of A that maximizes its advantage with respect to B.

Example 1:

Input: A = [2,7,11,15], B = [1,10,4,11]
Output: [2,11,7,15]


Example 2:

Input: A = [12,24,8,32], B = [13,25,32,11]
Output: [24,32,8,12]

*/

func advantageCount(A []int, B []int) []int {

	aSort := make([]int, len(A)) //Defino longitud para que el copy no me genere un slice vacio
	bSort := make([]int, len(B))
	copy(aSort, A)
	copy(bSort, B)
	//Ordeno de menor a mayor aSort(copia de A) u bSort(copia de B)
	sort.Sort(sort.IntSlice(aSort))
	sort.Sort(sort.IntSlice(bSort))

	//Mapas que guardan como clave,cada valor del array, y como valor del mapa, el índice del array
	mapHelperA := make(map[int]int, len(A))
	mapHelperB := make(map[int]int, len(B))

	//slice donde se irá appendeando el resultado
	aResult := make([]int, len(A))

	for i, v := range A {
		mapHelperA[v] = i
	}

	for i, v := range B {
		mapHelperB[v] = i
	}

	for _, v := range aSort {
		//comparo cada valor de aSort con el primer valor de bSort. Si es mayor, reposiciono el valor en A segun el indice que
		//le corresponda en B, y borro el elemento en bSort para que no vuelva a ser comparado
		valueB := bSort[0]
		if v > valueB {
			//Si el valor de A es mayor al de B, entonces cambio el valor de A para que quede a la misma altura del indice
			//original que le corresponde a ese valor en B
			destino := mapHelperB[valueB]
			permutacionA(destino, v, &aResult)
			//genero una copia de bSort sacandole el indice 0 que ya fue procesado
			bSort = bSort[1:]
		} else {
			//si el valor de A no es mayor al de B, lo dejo en el mismo indice donde esta
			destino := mapHelperA[v]
			permutacionA(destino, v, &aResult)
		}
	}
	return aResult
}

func permutacionA(indice int, valor int, aResult *[]int) {
	(*aResult)[indice] = valor
}

/*
A: {2, 0, 4, 1, 2},
B: {1, 3, 0, 0, 2},
result:{2, 0, 2, 1, 4},

result mio [2 0 0 2 4]
*/

// permutacionA(1, 3, &A)
// funcion que intercambia los valores de dos indices

// func permutacionB(origen int, destino int, A *[]int) {
// 	valueA := (*A)[origen]
// 	valueB := (*A)[destino]

// 	(*A)[origen] = valueB
// 	(*A)[destino] = valueA

// }

//Solution 1
/*

func advantageCount(A []int, B []int) []int {
	magic := [][2]int{}
	for i, v := range B {
		magic = append(magic, [2]int{i, v})
	}
	sort.Slice(magic, func(i, j int) bool {
		return magic[i][1] > magic[j][1]
	})
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	l, r := 0, len(A)-1
	C := make([]int, len(A))
	for _, v := range magic {
		if A[l] > v[1] {
			C[v[0]] = A[l]
			l++
		} else {
			C[v[0]] = A[r]
			r--
		}
	}
	return C
}

*/
