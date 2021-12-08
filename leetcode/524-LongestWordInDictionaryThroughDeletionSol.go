package leetcode

func solution524(s string, d []string) string {
	ans := ""
	for _, ss := range d {
		i, j := 0, 0
		for i < len(s) && j < len(ss) {
			if s[i] == ss[j] {
				i++
				j++
			} else {
				i++
			}
		}
		if j == len(ss) && len(ss) >= len(ans) {
			if len(ss) > len(ans) {
				ans = ss
			} else {
				if ss < ans {
					ans = ss
				}
			}
		}
	}
	return ans
}

//turns out this solution doesn't work for this case:
// s: ew, d: ["we", "w"]
// the output would be "we", but the answer should be "w".
func solution524_bad(s string, d []string) string {
	if len(s) == 0 {
		return ""
	}
	signature := getDic524(s)
	ans := ""
	for _, ss := range d {
		tmp := getDic524(ss)
		match := true
		for k, v := range tmp {
			if sv, ok := signature[k]; ok {
				if sv < v {
					match = false
					break
				}
			} else {
				match = false
			}
			if !match {
				break
			}
		}
		if match {
			if len(ss) > len(ans) {
				ans = ss
			} else if len(ss) == len(ans) {
				if ss < ans {
					ans = ss
				}
			}
		}
	}
	return ans
}

func getDic524(s string) map[rune]int {
	counter := make(map[rune]int)
	for _, c := range s {
		counter[c]++
	}
	return counter
}
