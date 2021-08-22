package dp

import (
	"fmt"
	"log"
)

/*
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？



示例 1:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	//This is how to make a 2-dimensional array(slice)
	memo := make([][]int, len(grid))
	for i := range memo {
		memo[i] = make([]int, len(grid[0]))
	}

	memo[0][0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		memo[0][i] = memo[0][i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		memo[i][0] = memo[i-1][0] + grid[i][0]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			memo[i][j] = max(memo[i-1][j], memo[i][j-1]) + grid[i][j]
		}
	}

	return memo[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Offer47 struct{}

type TestCaseOffer47 struct {
	grid [][]int
	ans  int
}

func (o *Offer47) Run() {
	tests := []TestCaseOffer47{
		{
			grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			ans:  12,
		},
		{
			grid: [][]int{{1, 2, 5}, {3, 2, 1}},
			ans:  9,
		},
	}
	for i, t := range tests {
		res := maxValue(t.grid)
		if res != t.ans {
			log.Fatalf("[Offer47] test %d failed, expect: %v, god: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer47] All tests passed")
}
