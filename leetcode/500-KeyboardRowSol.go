package leetcode

import (
	"strings"
)

func solution500(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}
	result := []string{}

	//build mapping
	mapping := map[byte]byte{}
	for _, c := range "qwertvuiop" {
		mapping[byte(c)] = byte('1')
	}
	for _, c := range "asdfghjkl" {
		mapping[byte(c)] = byte('2')
	}
	for _, c := range "zxcvbnm" {
		mapping[byte(c)] = byte('3')
	}
	for _, s := range input {
		if validForSol500(strings.ToLower(s), mapping) {
			result = append(result, s)
		}
	}
	return result
}

func validForSol500(s string, m map[byte]byte) bool {
	var prev byte
	for _, c := range s {
		if v, ok := m[byte(c)]; ok {
			if prev == 0 {
				prev = v
			} else {
				if prev != v {
					return false
				}
			}
		}
	}
	return true
}
