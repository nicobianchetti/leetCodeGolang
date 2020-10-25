package main

import "fmt"

func sherlockAndAnagrams(s string) int32 {
	isTiene := false
	alfa := make(map[int]int, 26)

	for _, v := range s {
		alfa[int(v)]++
		if alfa[int(v)] > 2 {
			isTiene = true
			break
		}

	}

	if isTiene == false {
		return 0
	}

	fmt.Print(alfa)

	return 1

}
