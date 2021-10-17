package leetcode

func solution417(heights [][]int) [][]int {
	res := [][]int{}
	if len(heights) == 0 {
		return res
	}
	//initialization
	rows, cols := len(heights), len(heights[0])
	atlantic, pacific := make([][]bool, rows), make([][]bool, rows)
	for i := 0; i < rows; i++ {
		atlantic[i] = make([]bool, cols)
		pacific[i] = make([]bool, cols)
	}
	dir := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	//dfs
	var dfs func([][]int, int, int, [][]bool)
	dfs = func(matrix [][]int, row, col int, visited [][]bool) {
		visited[row][col] = true
		for i := 0; i < len(dir); i++ {
			r := row + dir[i][0]
			c := col + dir[i][1]
			if r < 0 || r >= rows || c < 0 || c >= cols || visited[r][c] || heights[r][c] < heights[row][col] {
				continue
			}
			dfs(matrix, r, c, visited)
		}
	}

	for i := 0; i < rows; i++ {
		dfs(heights, i, 0, pacific)
		dfs(heights, i, cols-1, atlantic)
	}

	for j := 0; j < cols; j++ {
		dfs(heights, 0, j, pacific)
		dfs(heights, rows-1, j, atlantic)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
