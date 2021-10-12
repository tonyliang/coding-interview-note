package leetcode

/*
 Why do we only store the "remain" in the memo?
 Think about this:

 [whatever + and - combination                earliest] ...cur_index, remain = 1
 [whatever + and - combination           mid]...cur_index, remain = 5
 [whatever + and - combination latest]...cur_index, remain = 23

 we are hoping that when we're done with the earliest and get the answer,
 we have already computed remain = 5 and remain = 23 in the journey.
 Therefore, when mid and latest gets called again, they can simply use memo[5] and memo[23].
 Seems like it has nothing to do with the current index of nums.
*/
func solution416(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	N := len(nums)
	memo := make(map[int]bool)

	var dfs func(int, int) bool
	dfs = func(start, remain int) bool {
		if remain == 0 {
			return true
		}
		if remain < 0 || start == N {
			return false
		}
		if v, ok := memo[remain]; ok {
			return v
		}

		res := dfs(start+1, remain-nums[start]) || dfs(start+1, remain)
		memo[remain] = res
		return memo[remain]
	}
	dfs(0, target)
	return memo[target]
}

// func solution416(nums []int) bool {
// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	if sum%2 == 1 {
// 		return false
// 	}
// 	target := sum / 2
// 	N := len(nums)
// 	dp := make([][]bool, N+1)
// 	for i := 0; i < N+1; i++ {
// 		dp[i] = make([]bool, sum+1)
// 	}
// 	for i := 1; i <= N; i++ {
// 		dp[i][nums[i-1]] = true
// 	}
// 	for i := 1; i <= N; i++ {
// 		for j := nums[i-1]; j <= target; j++ {
// 			dp[i][j] = dp[i][j] || dp[i-1][j] || dp[i-1][j-nums[i-1]]
// 		}
// 	}
// 	return dp[N][target]
// }

//slow
func solution416_slow(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	N := len(nums)
	suffixSum := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		if i == N-1 {
			suffixSum[i] = nums[i]
		} else {
			suffixSum[i] = suffixSum[i+1] + nums[i]
		}
	}
	res := false
	var dfs func(int, int)
	dfs = func(start, cur int) {
		if start == N {
			if cur == target {
				res = true
			}
			return
		}
		if suffixSum[start] < target-cur {
			return
		}
		dfs(start+1, cur+nums[start])
		dfs(start+1, cur)
	}
	dfs(0, 0)
	return res
}
