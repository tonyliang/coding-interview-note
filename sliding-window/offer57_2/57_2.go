//https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
package offer57_2

import (
	"fmt"
	"log"
)

func findContinuousSequence(target int) [][]int {
	result := [][]int{}
	if target < 1 {
		return result
	}
	data := []int{}
	for i := 1; i < target; i++ {
		data = append(data, i)
	}

	total := 0
	right := 0
	left := 0

	for right < len(data) {
		total += data[right]
		for total > target {
			total -= data[left]
			left++
		}
		if total == target {
			result = append(result, data[left:right+1])
		}
		right++
	}

	return result
}

type TestCase struct {
	target int
	ans    [][]int
}

func equal2(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !equal(v, b[i]) {
			return false
		}
	}
	return true
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Run() {
	tests := []TestCase{
		{
			target: 9,
			ans:    [][]int{{2, 3, 4}, {4, 5}},
		},
		{
			target: 3,
			ans:    [][]int{{1, 2}},
		},
		{
			target: 15,
			ans:    [][]int{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}},
		},
	}
	for i, t := range tests {
		res := findContinuousSequence(t.target)
		if !equal2(res, t.ans) {
			log.Fatalf("Test %d failed. Expect: %v, Got: %v", i, t.ans, res)
		}
	}
	fmt.Println("All tests passed")
}
