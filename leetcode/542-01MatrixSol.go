package leetcode

type Cell struct {
	r int
	c int
}

func solution542(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	M, N := len(matrix), len(matrix[0])
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == 0 {
				matrix[i][j] = -1
			}
			if matrix[i][j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
	queue := make([]Cell, 0)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == -1 {
				queue = append(queue, Cell{i, j})
			}
		}
	}
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	level := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cel := queue[0]
			queue = queue[1:]
			for k := 0; k < 4; k++ {
				r, c := cel.r+dir[k][0], cel.c+dir[k][1]
				if r < 0 || r >= M || c < 0 || c >= N || matrix[r][c] != 0 {
					continue
				}
				matrix[r][c] = level
				queue = append(queue, Cell{r, c})
			}
		}
		level++
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}
