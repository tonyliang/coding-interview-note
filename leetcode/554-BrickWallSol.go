package leetcode

func solution554(wall [][]int) int {
	var width int
	seen := make(map[int]int)
	for i := 0; i < len(wall); i++ {
		cur := 0
		for j := 0; j < len(wall[i]); j++ {
			cur += wall[i][j]
			seen[cur]++
		}
		width = cur
	}

	maximum := len(wall)
	ans := maximum
	for k, v := range seen {
		if k > 0 && k < width {
			ans = min(ans, maximum-v)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
