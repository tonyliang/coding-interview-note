package leetcode

import "fmt"

var dir529 [][]int = [][]int{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

func solution529(board [][]byte, click []int) [][]byte {
	seen := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		seen[i] = make([]bool, len(board[0]))
	}
	var dfs func(int, int)
	dfs = func(r, c int) {
		seen[r][c] = true
		numMine := getNumMine(board, r, c)
		if numMine > 0 {
			b := []byte(fmt.Sprintf("%d", numMine))
			board[r][c] = b[0]
			return
		}
		board[r][c] = 'B'
		for _, d := range dir529 {
			newR, newC := r+d[0], c+d[1]
			if newR < 0 || newR >= len(board) || newC < 0 || newC >= len(board[0]) || board[newR][newC] == 'M' || seen[newR][newC] {
				continue
			}
			dfs(newR, newC)
		}
	}
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	dfs(click[0], click[1])
	return board
}

func getNumMine(board [][]byte, r, c int) int {
	count := 0
	for _, d := range dir529 {
		newR, newC := r+d[0], c+d[1]
		if newR < 0 || newR >= len(board) || newC < 0 || newC >= len(board[0]) {
			continue
		}
		if board[newR][newC] == 'M' {
			count++
		}
	}

	return count
}
