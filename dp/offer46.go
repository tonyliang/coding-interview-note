package dp

import (
	"fmt"
	"log"
	"strconv"
)

/*
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。



示例 1:

输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"


提示：

0 <= num < 231


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Okay
// func translateNum(num int) int {
// 	numStr := strconv.Itoa(num)
// 	total := 0
// 	translateNumHelper(numStr, 0, &total)
// 	return total
// }

// func translateNumHelper(num string, start int, total *int) {
// 	if start >= len(num) {
// 		*total++
// 		return
// 	}
// 	translateNumHelper(num, start+1, total)
// 	if start+2 <= len(num) && num[start] != '0' {
// 		v, err := strconv.Atoi(num[start : start+2])
// 		if err == nil && v < 26 {
// 			translateNumHelper(num, start+2, total)
// 		}
// 	}
// }

//Better (Memoization)
// func translateNum(num int) int {
// 	numStr := strconv.Itoa(num)
// 	memo := make(map[int]int)
// 	return translateNumHelper(numStr, 0, memo)
// }

// func translateNumHelper(num string, start int, memo map[int]int) int {
// 	if start == len(num) {
// 		return 1
// 	}
// 	if start > len(num) {
// 		return 0
// 	}
// 	if m, ok := memo[start]; ok {
// 		return m
// 	}
// 	res1 := translateNumHelper(num, start+1, memo)
// 	res2 := 0
// 	if num[start] != '0' && start+2 <= len(num) {
// 		v, err := strconv.Atoi(num[start : start+2])
// 		if err == nil && v < 26 {
// 			res2 = translateNumHelper(num, start+2, memo)
// 		}
// 	}
// 	memo[start] = res1 + res2
// 	return memo[start]
// }

//Best (O(n) time and O(n) space)
func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	N := len(numStr)
	dp := make([]int, 2)
	//str:       x1,  x2,  x3,  x4
	//dp:  dp0, dp1, dp2, dp3, dp4
	dp[0] = 1 //basecase (trick), if no char in str, total ways is 1.
	dp[1] = 1 //if one char in str only, total ways is 1.
	if N <= 1 {
		return 1
	}

	//starting from char 2, so that we start with str[0:2] i.e x1,x2
	c := 0
	for i := 2; i <= N; i++ {
		//if str[i-2: i] is a legit double-digit number (>=10 && <= 25)
		if legitNum(numStr[i-2 : i]) {
			//[ideally, but taking too many spaces] c = dp[i-2] + dp[i-1]
			c = dp[0] + dp[1]
		} else {
			c = dp[1]
		}
		//swap to save space
		dp[0] = dp[1]
		dp[1] = c
	}

	return dp[1]
}

func legitNum(s string) bool {
	if s[0] == '0' {
		return false
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return v >= 10 && v <= 25
}

type Offer46 struct{}

type TestCaseOffer46 struct {
	num int
	ans int
}

func (o *Offer46) Run() {
	tests := []TestCaseOffer46{
		{
			num: 1,
			ans: 1,
		},
		{
			num: 12,
			ans: 2,
		},
		{
			num: 123,
			ans: 3,
		},
		{
			num: 12258,
			ans: 5,
		},
		{
			num: 506,
			ans: 1,
		},
	}
	for i, t := range tests {
		res := translateNum(t.num)
		if res != t.ans {
			log.Fatalf("[Offer46] test %d failed, expect %v, got %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Offer46] All tests passed")
}
