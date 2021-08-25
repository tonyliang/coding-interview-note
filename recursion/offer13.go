package recursion

import (
	"fmt"
	"log"
)

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？



示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：

输入：m = 3, n = 1, k = 0
输出：1
提示：

1 <= n,m <= 100
0 <= k <= 20


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	movingCountHelper(0, 0, m, n, k, visited)

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] {
				ans++
			}
		}
	}
	return ans
}

func movingCountHelper(r, c, m, n, k int, visited [][]bool) {
	//when reaching here, (r,c) is a valid move.
	visited[r][c] = true
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, d := range dir {
		newR, newC := r+d[0], c+d[1]
		if newR >= 0 && newR < m && newC >= 0 && newC < n && !visited[newR][newC] && validMove(newR, newC, k) {
			movingCountHelper(newR, newC, m, n, k, visited)
		}
	}
}

func validMove(r, c, k int) bool {
	total := 0
	i := r
	for i > 0 {
		total += i % 10
		i = i / 10
	}
	i = c
	for i > 0 {
		total += i % 10
		i = i / 10
	}
	return total <= k
}

type Offer13 struct{}
type TestCaseOffer13 struct {
	m   int
	n   int
	k   int
	ans int
}

func (o *Offer13) Run() {
	tests := []TestCaseOffer13{
		{
			m:   3,
			n:   2,
			k:   1,
			ans: 3,
		},
		{
			m:   3,
			n:   1,
			k:   0,
			ans: 1,
		},
	}
	for i, t := range tests {
		res := movingCount(t.m, t.n, t.k)
		if res != t.ans {
			log.Fatalf("[Offer13] test %d failed, expect: %v, got: %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Offer13] All tests passed")
}
