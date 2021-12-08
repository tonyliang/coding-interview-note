package leetcode

func solution576(m, n, maxMove, startRow, startColumn int) int {
	//initialize memo: [m][n][maxMove+1]
	memo := make([][][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, maxMove+1)
			for k := 0; k <= maxMove; k++ {
				memo[i][j][k] = -1
			}
		}
	}
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var dfs func(int, int, int) int
	dfs = func(r, c, remain int) int {
		if r < 0 || c < 0 || r >= m || c >= n {
			return 1
		}
		if remain == 0 {
			memo[r][c][remain] = 0
			return 0
		}
		if memo[r][c][remain] > -1 {
			return memo[r][c][remain]
		}
		res := 0
		for i := 0; i < 4; i++ {
			newR, newC := r+dirs[i][0], c+dirs[i][1]
			res += dfs(newR, newC, remain-1) % 1000000007
		}
		memo[r][c][remain] = res % 1000000007
		return memo[r][c][remain]
	}
	return dfs(startRow, startColumn, maxMove)
}

func solution576_slow(m, n, maxMove, startRow, startColumn int) int {
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var ans int
	var dfs func([][]bool, int, int, int)
	dfs = func(visited [][]bool, remain, r, c int) {
		if remain == 0 {
			return
		}
		for _, d := range dir {
			newR, newC := r+d[0], c+d[1]
			if newR < 0 || newR >= m || newC < 0 || newC >= n {
				ans = (ans + 1) % 1000000007
				continue
			}

			dfs(visited, remain-1, newR, newC)

		}
	}
	dfs(visited, maxMove, startRow, startColumn)
	return ans % 1000000007
}
