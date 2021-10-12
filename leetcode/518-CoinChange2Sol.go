package leetcode

//better solution
func solution518(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

//okay solution
func solution518_ok(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	memo := make([][]int, amount+1)
	for i := 0; i < amount+1; i++ {
		memo[i] = make([]int, len(coins))
		for j := 0; j < len(coins); j++ {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(a, start int) int {
		if a == 0 {
			return 1
		}
		if memo[a][start] > -1 {
			return memo[a][start]
		}
		res := 0
		for i := start; i < len(coins); i++ {
			if a >= coins[i] {
				res += dfs(a-coins[i], i)
			}
		}
		memo[a][start] = res
		return res
	}
	dfs(amount, 0)
	return memo[amount][0]
}

//this solution will count duplicate cases
//ex: 1+2, 2+1 should be the same, but actually counted twice!
func solution518_dup(amount int, coins []int) int {
	if amount == 0 {
		return 0
	}
	memo := make(map[int]int)
	var dfs func(int) int

	dfs = func(a int) int {
		if a == 0 {
			return 1
		}
		if a < 0 {
			return 0
		}
		if v, ok := memo[a]; ok {
			return v
		}
		res := 0
		for _, c := range coins {
			if c <= a {
				res += dfs(a - c)
			}
		}
		memo[a] = res
		return memo[a]
	}

	dfs(amount)

	return memo[amount]
}
