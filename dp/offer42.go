package dp

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

/*
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。



示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curMax := -helper.MaxInt
	result := -helper.MaxInt
	for _, v := range nums {
		curMax = findMax(curMax+v, v)
		result = findMax(curMax, result)
	}
	return result
}

func findMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type TestCaseOffer42 struct {
	nums []int
	ans  int
}

type Offer42 struct{}

func (o *Offer42) Run() {
	tests := []TestCaseOffer42{
		{
			nums: []int{},
			ans:  0,
		},
		{
			nums: []int{1},
			ans:  1,
		},
		{
			nums: []int{1, 2, 3},
			ans:  6,
		},
		{
			nums: []int{-1, -2, -3},
			ans:  -1,
		},
		{
			nums: []int{-1, -2, 3},
			ans:  3,
		},
		{
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			ans:  6,
		},
	}
	for i, t := range tests {
		res := maxSubArray(t.nums)
		if res != t.ans {
			log.Fatalf("[Offer42] test %d failed, expect %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer42] All tests passed")
}
