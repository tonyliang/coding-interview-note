package leetcode

import "math"

func solution507(num int) bool {
	if num <= 1 {
		return false
	}
	sum, bound := 1, int(math.Sqrt(float64(num)))+1
	for i := 2; i < bound; i++ {
		if num%i != 0 {
			continue
		}
		corrDiv := num / i
		sum += corrDiv + i
	}
	return sum == num
}

//slow
func solution507_2(num int) bool {
	cur := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			cur += i
		}
	}
	return cur == num
}
