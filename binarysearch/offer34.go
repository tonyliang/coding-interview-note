package binarysearch

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？


示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	result := []int{-1, -1}
	result[0] = searchRangeLeftHelper(nums, target)  //O(logN)
	result[1] = searchRangeRightHelper(nums, target) //O(logN)
	//total: O(2LogN) = O(LogN)
	return result
}

func searchRangeRightHelper(nums []int, target int) int {
	//we want to maintain left + 1 < right in the loop
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}
	//when we exit the loop, left + 1 = right
	if nums[right] == target {
		return right
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func searchRangeLeftHelper(nums []int, target int) int {
	//we want to maintain left < right in the loop
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	//when we exit the loop, left = right
	if nums[left] == target {
		return left
	}
	return -1
}

type Offer34 struct{}
type TestCaseOffer34 struct {
	nums   []int
	target int
	ans    []int
}

func (o *Offer34) Run() {
	tests := []TestCaseOffer34{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			ans:    []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			ans:    []int{-1, -1},
		},
		{
			nums:   []int{},
			target: 0,
			ans:    []int{-1, -1},
		},
		{
			nums:   []int{8, 8, 8},
			target: 8,
			ans:    []int{0, 2},
		},
		{
			nums:   []int{8},
			target: 8,
			ans:    []int{0, 0},
		},
	}
	for i, t := range tests {
		res := searchRange(t.nums, t.target)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Offer34] test %d failed, expect: %v, got: %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Offer34] All tests passed")
}
