package main

import (
	"reflect"
	"sort"
)

/*

https://leetcode.com/problems/valid-anagram/

Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

*/

func isAnagram(s string, t string) bool {
	var sortS []int
	var sortT []int

	for _, v := range s {
		sortS = append(sortS, int(v))
	}

	for _, v := range t {
		sortT = append(sortT, int(v))
	}

	sort.Ints(sortS)
	sort.Ints(sortT)

	return reflect.DeepEqual(sortS, sortT)
}
