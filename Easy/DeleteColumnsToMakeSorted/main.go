package main

//https://leetcode.com/problems/delete-columns-to-make-sorted/

/*
Descripción: armo una cadena de caracteres por cada indice del array. Si la cadena formada es decreciente,
entonces sumo 1 al contador y elimino ese indice hasta que me queden todas cadenas con caracteres ordenados de
menor a mayor
*/

func minDeletionSize(A []string) int {
	var countAux int
	for j := 0; j < len(A[0]); j++ {
		for i := 1; i < len(A); i++ {
			if A[i][j] < A[i-1][j] {
				countAux++
				break
			}
		}
	}
	return countAux
}

/*
Solución 1:
Armo un mapa helper donde guardo ordenadas las cadenas de caracteres de cada indice
*/

// func minDeletionSize(A []string) int {
// 	var mapHelper = make(map[int]string)
// 	for _, value := range A {
// 		for index, valueChar := range value {
// 			mapHelper[index] += string(valueChar)
// 		}
// 	}

// 	var countIndex int
// 	for _, value := range mapHelper {
// 		var countAux int
// 		for _, valueChar := range value {
// 			if int(valueChar) >= countAux {
// 				countAux = int(valueChar)
// 			} else {
// 				countIndex++
// 				break
// 			}
// 		}
// 	}
// 	return countIndex
// }

/*
 Solución 2 : se agrega uso de goroutines a solución  1
*/

// var CountIndex int

// func Procesar(c chan string, done chan struct{}) {

// 	for element := range c {
// 		var countAux int
// 		for _, valueChar := range element {
// 			if int(valueChar) >= countAux {
// 				countAux = int(valueChar)
// 			} else {
// 				CountIndex++
// 				break
// 			}
// 		}
// 	}

// 	done <- struct{}{}
// 	close(done)
// }

// func minDeletionSize(A []string) int {
// 	CountIndex = 0
// 	var mapHelper = make(map[int]string)
// 	for _, value := range A {
// 		for index, valueChar := range value {
// 			mapHelper[index] += string(valueChar)
// 		}
// 	}

// 	sizeChanel := len(mapHelper)
// 	chanValues := make(chan string, sizeChanel)
// 	done := make(chan struct{})

// 	go Procesar(chanValues, done)

// 	for _, value := range mapHelper {
// 		chanValues <- value
// 	}

// 	close(chanValues)

// 	<-done

// 	return CountIndex
// }
