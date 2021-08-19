//https://leetcode.com/problems/permutations/
package permutation

import (
	"fmt"
	"log"
)

var result [][]int

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	helper(0, nums)
	return result
}

func helper(index int, nums []int) {
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		result = append(result, tmp)
		return
	}
	for i := index; i < len(nums); i++ {
		swap(nums, i, index)
		helper(index+1, nums)
		swap(nums, i, index)
	}
}

func swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}

type TestCase struct {
	nums []int
	ans  [][]int
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
			nums: []int{},
			ans:  [][]int{},
		},
		{
			nums: []int{1},
			ans:  [][]int{{1}},
		},
		{
			nums: []int{1, 2},
			ans:  [][]int{{1, 2}, {2, 1}},
		},
		{
			nums: []int{0, 1},
			ans:  [][]int{{0, 1}, {1, 0}},
		},
	}
	for i, t := range tests {
		res := permute(t.nums)
		if !equal2(res, t.ans) {
			log.Fatalf("Test %d failed. Expect %v, Got: %v\n", i, t.ans, res)
		}
	}
	fmt.Println("All tests passed")
}
