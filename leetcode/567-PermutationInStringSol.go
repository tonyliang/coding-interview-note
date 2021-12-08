package leetcode

func solution567(s1 string, s2 string) bool {
	if len(s1) == 0 {
		return false
	}
	target := make(map[byte]int)
	for _, c := range s1 {
		target[byte(c)]++
	}
	targetLen := len(target)
	curLen := 0
	cur := make(map[byte]int)
	var left, right int
	for right < len(s2) {
		if v, ok := target[s2[right]]; ok {
			cur[s2[right]]++
			if cur[s2[right]] == v {
				curLen++
			}
			if curLen == targetLen {
				return true
			}
			for cur[s2[right]] > v {
				if cur[s2[left]] == target[s2[left]] {
					curLen--
				}
				cur[s2[left]]--
				left++
			}

		} else {
			cur = make(map[byte]int)
			curLen = 0
			left = right + 1
		}
		right++
	}
	return false
}
