package leetcode

import "math/rand"

type Solution528 struct {
	prefixSum []int
}

func Constructor528(w []int) Solution528 {
	prefixSum := make([]int, len(w))
	for i := 0; i < len(w); i++ {
		if i == 0 {
			prefixSum[i] = w[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + w[i]
		}
	}
	sol := Solution528{prefixSum: prefixSum}
	return sol
}

func (s *Solution528) PickIndex() int {
	N := len(s.prefixSum)
	target := rand.Intn(s.prefixSum[N-1]) + 1
	lo, hi := 0, N-1
	for lo < hi {
		mid := lo + int((hi-lo)/2)
		if s.prefixSum[mid] == target {
			return mid
		} else if s.prefixSum[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor528(w);
 * param_1 := obj.PickIndex();
 */
