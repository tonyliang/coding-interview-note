package p154

import (
	"fmt"
	"log"
)

//https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,4,4,5,6,7] might become:

[4,5,6,7,0,1,4] if it was rotated 4 times.
[0,1,4,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums that may contain duplicates, return the minimum element of this array.


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//
//Case 1, [3,4,5,1,2], mid is "5" => mid(5) > right(2) => left = mid + 1
//Case 2, [3,1,3], mid is "1" => mid(1) < right(3) => right = mid
//Case 3, [3,3,3], mid is "3" => right--
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	//left == right now
	return nums[left]
}

type TestCase struct {
	nums []int
	ans  int
}

func Run() {
	tests := []TestCase{
		{
			nums: []int{3, 4, 5, 1, 2},
			ans:  1,
		},
		{
			nums: []int{3, 1, 3},
			ans:  1,
		},
		{
			nums: []int{3, 3, 3},
			ans:  3,
		},
		{
			nums: []int{2, 2, 2, 0, 1},
			ans:  0,
		},
		{
			nums: []int{1, 3, 5},
			ans:  1,
		},
	}

	for i, t := range tests {
		res := findMin(t.nums)
		if res != t.ans {
			log.Fatalf("Test %d failed. Expect %v, Got %v", i, t.ans, res)
		}
	}
	fmt.Println("All tests passed")
}
