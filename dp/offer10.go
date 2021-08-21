package dp

import (
	"fmt"
	"log"
)

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	a := 1
	b := 1
	var c int
	for i := 2; i <= n; i++ {
		c = (a + b) % (1e9 + 7)
		a = b
		b = c
	}
	return c
}

type TestCase1 struct {
	n   int
	ans int
}

func Run1() {
	tests := []TestCase1{
		{
			n:   0,
			ans: 1,
		},
		{
			n:   1,
			ans: 1,
		},
		{
			n:   2,
			ans: 2,
		},
		{
			n:   7,
			ans: 21,
		},
	}
	for i, t := range tests {
		res := numWays(t.n)
		if res != t.ans {
			log.Fatalf("Test %d failed. Expect: %d, Got %d\n", i, t.ans, res)
		}
	}
	fmt.Println("All tests passed")
}
