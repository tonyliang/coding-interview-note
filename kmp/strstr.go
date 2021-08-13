package kmp

import (
	"fmt"
)

//Brute Force Search O(mn)
func BFSearch(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	i, j := 0, 0
	for i < l1 && j < l2 {
		if haystack[i] == needle[j] { //matching and advancing
			i++
			j++
		} else {
			//i.e i = i - j + 1
			//to put i at the previous state (before the matching and advancing) plus 1.
			//[note] when j = 0, i = i - 0 + 1, which advance i by one, with which infinite loop wont happen.
			i -= j - 1
			j = 0
		}
	}
	if j == l2 {
		//i-j is the beginning index of the matching and advancing process.
		return i - j
	}
	return -1
}

//https://www.youtube.com/watch?v=GTJr8OvyEVQ&t=760s
//O(m + n)
func KMPSearch(haystack string, needle string) int {
	lps := makeNextTable(needle)
	l1 := len(haystack)
	l2 := len(needle)
	i, j := 0, 0
	for i < l1 && j < l2 {
		if haystack[i] == needle[j] { //matching and advancing
			i++
			j++
		} else {
			if j > 0 {
				j = lps[j-1]
			} else {
				//this is to avoid infinite loop.
				//because when (j == 0 and haystack[i] != needle[j]), we need to advance i to break the loop.
				i++
			}
		}
	}
	if j == l2 {
		//i-j is the beginning index of the matching and advancing process.
		return i - j
	}
	return -1
}

func makeNextTable(pattern string) []int {
	nextTable := make([]int, len(pattern))
	j := 0
	i := 1
	for i < len(nextTable) {
		if pattern[i] == pattern[j] {
			nextTable[i] = j + 1
			j++
			i++
		} else {
			if j != 0 {
				j = nextTable[j-1]
			} else {
				nextTable[i] = 0
				i++
			}
		}
	}
	return nextTable
}

func Run() {
	fmt.Println(BFSearch("ABABCDE", "ABCD"))  //expect: 2
	fmt.Println(KMPSearch("ABABCDE", "ABCD")) //expect: 2
}
