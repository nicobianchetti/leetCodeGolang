package main

//https://leetcode.com/problems/delete-columns-to-make-sorted/

func minDeletionSize(A []string) int {
	var mapHelper = make(map[int]string)
	for _, value := range A {
		for index, valueChar := range value {
			mapHelper[index] += string(valueChar)
		}
	}
	// log.Println(mapHelper)

	var countIndex int
	for _, value := range mapHelper {
		var countAux int
		for _, valueChar := range value {
			if int(valueChar) > countAux {
				countAux = int(valueChar)
				// fmt.Println("Leyendo indice", index, "char", string(valueChar))
			} else {
				countIndex++
				// fmt.Println("Entra aca leyendo indice", index, "char", string(valueChar))
				break
			}
		}
	}
	return countIndex
}
