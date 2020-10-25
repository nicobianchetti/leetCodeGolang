package main

import (
	"strconv"
)

/*

https://leetcode.com/problems/binary-gap/

Given a positive integer N, find and return the longest distance between two consecutive 1's in the binary representation of N.

If there aren't two consecutive 1's, return 0.

Example 1:

Input: 22
Output: 2
Explanation:
22 in binary is 0b10110.
In the binary representation of 22, there are three ones, and two consecutive pairs of 1's.
The first consecutive pair of 1's have distance 2.
The second consecutive pair of 1's have distance 1.
The answer is the largest of these two distances, which is 2.

---------------------------------------

Example 2:

Input: 5
Output: 2
Explanation:
5 in binary is 0b101.

---------------------------------------


Example 3:

Input: 6
Output: 1
Explanation:
6 in binary is 0b110.

---------------------------------------


Example 4:

Input: 8
Output: 0
Explanation:
8 in binary is 0b1000.
There aren't any consecutive pairs of 1's in the binary representation of 8, so we return 0.
*/

func binaryGap(N int) int {
	//Convierto el int de entrada en un binario,ésto me devuelve una cadena de bytes que puedo recorrrer
	Binary := strconv.FormatInt(int64(N), 2)

	Helper := make([]int, len(Binary))

	var countMax int
	var countAux int
	//recorro la cadena de byes. Para traspasarlos a un array de int, primero convierto el byte(Ej: 49) a string("1")
	//Luego el string lo covierto a int(1) usando strconv.Atoi
	for i, v := range Binary {
		Helper[i], _ = strconv.Atoi(string(v))
	}

	isPrimerUno := true

	for _, v := range Helper {
		if v == 1 {
			if isPrimerUno {
				isPrimerUno = false
			} else {
				countAux++
				if countAux > countMax {
					countMax = countAux
				}
				countAux = 0
			}
		} else {
			if !isPrimerUno {
				countAux++
			}
		}
	}

	return countMax
}
