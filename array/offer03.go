package array

/*
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3


限制：

2 <= n <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//Even better
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
		} else {
			//because nums[i] != i , the correct index of nums[i] is nums[i],
			//before we move nums[i] to its correct index, we check if nums[nums[i]] is nums[i],
			//if yes, duplicate right there, we can stop and return the answer.
			//if not, we move nums[i] to its correct index by swapping values.
			if nums[nums[i]] == nums[i] {
				return nums[i]
			} else {
				//swap (moving nums[i] to its correct index
				//might be in a loop until it gets swapped to the right location
				SwapInt(nums, i, nums[i])
			}
		}
	}
	return -1
}

//Add 1 to each element so as to handle 0 value at index 0.
//also this modifies the original array.
func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		nums[i]++
	}
	for _, n := range nums {
		if nums[abs(n)-1] < 0 {
			return abs(n) - 1
		} else {
			nums[abs(n)-1] *= -1
		}
	}
	return -1
}
