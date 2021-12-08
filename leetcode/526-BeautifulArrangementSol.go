package leetcode

func solution526(N int) int {
	ans := 0
	var dfs func([]bool, int)

	dfs = func(seen []bool, pos int) {

		if pos == N {
			ans++
			return
		}
		for i := 1; i <= N; i++ {

			if seen[i-1] {
				continue
			}

			if i%(pos+1) == 0 || (pos+1)%i == 0 {
				seen[i-1] = true
				dfs(seen, pos+1)
				seen[i-1] = false
			}
		}
	}
	dfs(make([]bool, N), 0)
	return ans
}
