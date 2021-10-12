package leetcode

func solution494(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	//speed up by making a suffix sum
	N := len(nums)
	suffixSum := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		if i == N-1 {
			suffixSum[i] = nums[i]
		} else {
			suffixSum[i] = suffixSum[i+1] + nums[i]
		}
	}

	res := 0
	var dfs func(int, int)

	dfs = func(start, cur int) {
		if start == len(nums) {
			if cur == target {
				res++
			}
			return
		}

		//if it's not enough to make it to target, stop
		if suffixSum[start] < target-cur {
			return
		}

		dfs(start+1, cur+nums[start])
		dfs(start+1, cur-nums[start])
	}

	dfs(0, 0)
	return res
}
