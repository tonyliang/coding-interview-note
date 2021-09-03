package dp

import (
	"fmt"
	"log"
)

/*
Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.

Return the maximum product you can get.



Example 1:

Input: n = 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.
Example 2:

Input: n = 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.


Constraints:

2 <= n <= 58

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/* slow: O(n ^ n)
func integerBreak(n int) int {
	ans := 0
	var dfs func(int, int, int)
	dfs = func(m, count, cur int) {
		if m == 0 {
			if count > 1 {
				ans = max(ans, cur)
			}
			return
		}
		for i := 1; i <= m; i++ {
			dfs(m-i, count+1, cur*i)
		}
	}
	dfs(n, 0, 1)
	return ans
}
*/

//time: O(N), memory: O(N)
func integerBreak(n int) int {
	memo := make(map[int]int)
	var dfs func(int, int) int
	dfs = func(m, count int) int {
		if m == 0 {
			if count > 1 {
				return 1
			}
			return 0
		}
		if v, ok := memo[m]; ok {
			if count > 1 {
				return max(v, m)
			}
			return v
		}
		r := 0
		for i := 1; i <= m; i++ {
			r = max(r, i*dfs(m-i, count+1))
		}
		memo[m] = r
		return r
	}
	return dfs(n, 0)
}

type P343 struct{}
type TestCaseP343 struct {
	n   int
	ans int
}

func (p *P343) Run() {
	tests := []TestCaseP343{
		{
			n:   2,
			ans: 1,
		},
		{
			n:   4,
			ans: 4,
		},
		{
			n:   10,
			ans: 36,
		},
		{
			n:   33,
			ans: 177147,
		},
		{
			n:   44,
			ans: 9565938,
		},
	}
	for i, t := range tests {
		res := integerBreak(t.n)
		if res != t.ans {
			log.Fatalf("[P343] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[P343] All tests passed")
}
