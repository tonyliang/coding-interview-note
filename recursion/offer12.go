package recursion

import (
	"fmt"
	"log"
)

/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。


示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
时间复杂度 O(3^K MN):  board size: MN, each helper run 3 ways, each way's depth is K (word length)
Space complexity O(MN) + O(K):  O(MN) is the bool matrix, O(K) is the call stack.
[NOTE] we can reduce space to O(K) by modifying the original matrix for visited.
*/

func exist(board [][]byte, word string) bool {
	if len(word) == 0 || board == nil {
		return false
	}
	R, C := len(board), len(board[0])
	visited := make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if board[i][j] == word[0] {

				res := offer12Helper(board, i, j, word, 0, visited)
				if res {
					return true
				}
			}
		}
	}
	return false
}

func offer12Helper(board [][]byte, r, c int, word string, k int, visited [][]bool) bool {
	//note: when reaching here, we have a match at kth byte of word!
	if k == len(word)-1 {
		return true
	}
	visited[r][c] = true

	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for i := 0; i < len(dir); i++ {
		newR, newC := r+dir[i][0], c+dir[i][1]
		if newR >= 0 && newR < len(board) && newC >= 0 && newC < len(board[0]) && !visited[newR][newC] && board[newR][newC] == word[k+1] {
			res := offer12Helper(board, newR, newC, word, k+1, visited)
			if res {
				return true
			}
		}
	}

	visited[r][c] = false
	return false
}

type Offer12 struct{}
type TestCaseOffer12 struct {
	board [][]byte
	word  string
	ans   bool
}

func (o *Offer12) Run() {
	tests := []TestCaseOffer12{
		{
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCCED",
			ans:   true,
		},
		{
			board: [][]byte{{'a', 'b'}, {'c', 'd'}},
			word:  "abcd",
			ans:   false,
		},
	}
	for i, t := range tests {
		res := exist(t.board, t.word)
		if res != t.ans {
			log.Fatalf("[Offer12] test %d failed, expect: %v, got: %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Offer12] All tests passed")
}
