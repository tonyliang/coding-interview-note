package p239

import (
	"fmt"
	"log"
)

type TestCase struct {
	nums []int
	k    int
	ans  []int
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		return make([]int, 0)
	}
	//[note] interesting pattern! there will be (len(nums) - k + 1) windows.
	result := make([]int, len(nums)-k+1)
	curIdx := 0

	queue := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		//enqueue the latest element
		queue = append(queue, nums[i])

		//if the max element in queue is outside the current window, remove it.
		if i-k >= 0 && nums[i-k] == queue[0] {
			queue = queue[1:]
		}

		//if a window is formed, output value
		if i-k >= -1 {
			result[curIdx] = queue[0]
			curIdx++
		}
	}
	return result
}

func Run() {
	tests := []TestCase{
		{
			nums: []int{1},
			k:    1,
			ans:  []int{1},
		},
		{
			nums: []int{1, -1},
			k:    1,
			ans:  []int{1, -1},
		},
		{
			nums: []int{9, 11},
			k:    2,
			ans:  []int{11},
		},
		{
			nums: []int{4, -2},
			k:    2,
			ans:  []int{4},
		},
		{
			nums: []int{1, 2, 3},
			k:    2,
			ans:  []int{2, 3},
		},
		{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			ans:  []int{3, 3, 5, 5, 6, 7},
		},
		{
			nums: []int{1, 2, 3},
			k:    2,
			ans:  []int{2, 3},
		},
	}
	for i, t := range tests {
		res := maxSlidingWindow(t.nums, t.k)
		if !equal(res, t.ans) {
			log.Fatalf("test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("All tests passed.")
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
