package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

func MaxArrayCase1(A []int) int {

	//Busco recorriendo todo el array
	start := time.Now()
	fmt.Println("Start es de tipo", reflect.TypeOf(start))
	var max int
	for _, v := range A {
		if v > max {
			max = v
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Done in %s /n", elapsed)

	return max
}

func MaxArrayCase2(A []int) int {
	//Busco haciendo sort descendiente
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	max := A[0]

	return max
}

func MaxArrayCase3(A []int) int {
	//Busco haciendo sort ascendiente
	sort.Sort(sort.IntSlice(A))
	maxIndex := len(A) - 1
	max := A[maxIndex]

	return max
}

// func TimeElapsed(start)
