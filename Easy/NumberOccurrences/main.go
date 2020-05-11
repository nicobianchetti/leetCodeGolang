package main

/*
https://leetcode.com/problems/unique-number-of-occurrences/

Runtime: 0 ms, faster than 100.00% .
Memory Usage: 2.3 MB, less than 100.00% .
*/

func uniqueOccurrences(arr []int) bool {
	var mapHelper = make(map[int]int)
	for _, v := range arr {
		_, found := mapHelper[v]
		if found {
			mapHelper[v]++
		} else {
			mapHelper[v] = 1
		}
	}

	isUnique := UniqueValue(mapHelper)

	return isUnique
}

func UniqueValue(mapHelper map[int]int) bool {
	for key, v := range mapHelper {
		for keyD, vD := range mapHelper {
			if keyD != key {
				if vD == v {
					return false
				}
			}
		}
	}
	return true
}
