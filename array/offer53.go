package array

import (
	"fmt"
	"log"
)

/*
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

示例 1:

输入: [0,1,3]
输出: 2
示例 2:

输入: [0,1,2,3,4,5,6,7,9]
输出: 8


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// When you see "sorted", think about binary search!
// If the length of nums is N, there are N elements in the array, and
// each element's value ranges from 0 to N.
// So if nums = [1,2,3,4], len = 4, each element value should ranges from 0 to 4.
// But one is missing, so you might get [0,1,2,4], which means 3 is missing.
//[Alternative thinking]: The size of nums is 4. In a perfect world, we should see
//[1,2,3,4] in nums, however, we allow 0 to be in the array, so one element has to be
//kicked out. Therefore this is possible: [0,1,2,3] or [0,1,2,4]
func missingNumber(nums []int) int {
	N := len(nums)
	//Note the value of each element in nums is in the range [0, N]
	left := 0
	right := N - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	//left == right now
	if nums[left] != left {
		return left
	} else {
		return left + 1
	}
}

type Offer53 struct{}
type TestCaseOffer53 struct {
	nums []int
	ans  int
}

func (o *Offer53) Run() {
	tests := []TestCaseOffer53{
		{
			nums: []int{0},
			ans:  1,
		},
		{
			nums: []int{1},
			ans:  0,
		},
		{
			nums: []int{0, 1, 2, 3},
			ans:  4,
		},
		{
			nums: []int{1, 2, 3, 4},
			ans:  0,
		},
		{
			nums: []int{0, 1, 3},
			ans:  2,
		},
		{
			nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 9},
			ans:  8,
		},
	}
	for i, t := range tests {
		res := missingNumber(t.nums)
		if res != t.ans {
			log.Fatalf("[Offer53] Test %d failed. Expect %d, Got %d", i, t.ans, res)
		}
	}
	fmt.Println("[Offer53] All tests passed")
}
