package matrix

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


限制：

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	seen := make([][]bool, len(matrix))
	for i := 0; i < len(seen); i++ {
		seen[i] = make([]bool, len(matrix[0]))
	}
	r, c, d := 0, 0, 1
	lenR, lenC := len(matrix), len(matrix[0])
	result := make([]int, lenR*lenC)
	for i := 0; i < len(result); i++ {
		result[i] = matrix[r][c]
		seen[r][c] = true
		_r, _c := r+dir[d][0], c+dir[d][1]
		if _r < 0 || _r >= lenR || _c < 0 || _c >= lenC || seen[_r][_c] {
			d = (d + 1) % 4
			_r, _c = r+dir[d][0], c+dir[d][1]
		}
		r, c = _r, _c
	}
	return result
}

type Offer29 struct{}
type TestCaseOffer29 struct {
	matrix [][]int
	ans    []int
}

func (o *Offer29) Run() {
	tests := []TestCaseOffer29{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			ans:    []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			ans:    []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			matrix: [][]int{},
			ans:    []int{},
		},
		{
			matrix: [][]int{{3}, {2}},
			ans:    []int{3, 2},
		},
	}
	for i, t := range tests {
		res := spiralOrder(t.matrix)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Offer29] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer29] All tests passed")
}
