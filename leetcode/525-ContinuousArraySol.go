package leetcode

func solution525(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	seen := make(map[int]int)

	//trick, so that we can handle answer that includes first element of array.
	seen[0] = -1

	cur := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cur += -1
		} else {
			cur += 1
		}
		target := cur
		if v, ok := seen[target]; ok {

			tmp := i - v
			if tmp > ans {
				ans = tmp
			}
		}

		if _, ok := seen[cur]; !ok {
			seen[cur] = i
		}

	}
	return ans
}
