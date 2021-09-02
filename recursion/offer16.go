package recursion

import (
	"fmt"
	"log"
)

/*
 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。



示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow(1/x, -n)
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	} else {
		return x * myPow(x*x, (n-1)/2)
	}
}

type Offer16 struct {
}

type TestCaseOffer16 struct {
	x   float64
	n   int
	ans float64
}

func (o *Offer16) Run() {
	tests := []TestCaseOffer16{
		{
			x:   2,
			n:   10,
			ans: 1024,
		},
		{
			x:   2.1,
			n:   3,
			ans: 9.261000000000001,
		},
		{
			x:   2,
			n:   -2,
			ans: 0.25,
		},
	}
	for i, t := range tests {
		res := myPow(t.x, t.n)
		if res != t.ans {
			log.Fatalf("[Offer16] test %d failed, expect: %v, got: %v", i, t.ans, res)

		}
	}
	fmt.Println("[Offer16] All tests passed")
}
