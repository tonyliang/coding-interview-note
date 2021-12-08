package leetcode

func solution547(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	total := len(M)
	parents := make([]int, total)
	ranks := make([]int, total)

	for i, _ := range parents {
		parents[i] = i
	}
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			if M[i][j] == 1 {
				union(i, j, parents, ranks, &total)
			}
		}
	}
	return total
}

func union(a, b int, parents, ranks []int, total *int) {
	pa, pb := find(a, parents), find(b, parents)
	if pa == pb {
		return
	}
	*total--
	if ranks[pa] > ranks[pb] {
		parents[pb] = pa
	} else {
		if ranks[pa] == ranks[pb] {
			ranks[pb]++
		}
		parents[pa] = pb
	}
}

func find(a int, parents []int) int {
	if parents[a] == a {
		return a
	}
	parents[a] = find(parents[a], parents)
	return parents[a]
}
