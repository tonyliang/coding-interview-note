package leetcode

func solution581(nums []int) int {
	//get the starting and ending points where array is not in ascending order.
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 0
	}
	var i int
	for i < len(nums)-1 {
		if nums[i] > nums[i+1] {
			break
		}
		i++
	}
	var j = len(nums) - 1
	for j > 0 {
		if nums[j] < nums[j-1] {
			break
		}
		j--
	}
	if i > j {
		return 0
	}

	//get the min and max values within the subarray
	tmpMin, tmpMax := getMin(nums[i:j+1]), getMax(nums[i:j+1])
	//fmt.Printf("%v, %v\n", tmpMin, tmpMax)
	//expand the left side of subarray until there is no values on the left side of the subarray that is greater than min
	for i > 0 && nums[i-1] > tmpMin {
		i--
	}
	//expand the right hand side of surbarray until there is no values on the right hand side of the subarray that is smaller than the max.
	for j < len(nums)-1 && nums[j+1] < tmpMax {
		j++
	}

	return j - i + 1
}

func getMin(arr []int) int {
	ans := 1<<(32-1) - 1
	for _, a := range arr {
		ans = min581(ans, a)
	}
	return ans
}

func getMax(arr []int) int {
	ans := -1 << (32 - 1)
	for _, a := range arr {
		ans = max581(ans, a)
	}
	return ans
}

func min581(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max581(a, b int) int {
	if a < b {
		return b
	}
	return a
}
