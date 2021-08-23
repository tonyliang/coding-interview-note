package strings

import (
	"fmt"
	"log"
)

/*
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。



示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	seen := make(map[byte]int)
	ans := 0
	left := 0
	right := 0
	for right < len(s) {
		seen[s[right]]++
		for seen[s[right]] > 1 {
			seen[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Offer48 struct{}

type TestCaseOffer48 struct {
	s   string
	ans int
}

func (o *Offer48) Run() {
	tests := []TestCaseOffer48{
		{
			s:   "abcabcbb",
			ans: 3,
		},
		{
			s:   "bbbbb",
			ans: 1,
		},
		{
			s:   "pwwkew",
			ans: 3,
		},
		{
			s:   "",
			ans: 0,
		},
	}
	for i, t := range tests {
		res := lengthOfLongestSubstring(t.s)
		if res != t.ans {
			log.Fatalf("[Offer48] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer48] All tests passed")
}
