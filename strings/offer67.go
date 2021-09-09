package strings

import (
	"fmt"
	"log"
	"math"
)

/*
写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。



首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。

当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：

假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func strToInt(str string) int {
	num := 0
	sign := 1
	i := 0
	n := len(str)
	for i < n && str[i] == ' ' {
		i++
	}
	start := i
	for ; i < n; i++ {
		if i == start && str[i] == '+' {
			sign = 1
		} else if i == start && str[i] == '-' {
			sign = -1
		} else if str[i] >= '0' && str[i] <= '9' {
			tmp := int(str[i] - '0')
			if num < math.MinInt32/10 || num == math.MinInt32/10 && -tmp < math.MinInt32%10 {
				return math.MinInt32
			}
			if num > math.MaxInt32/10 || num == math.MaxInt32/10 && tmp > math.MaxInt32%10 {
				return math.MaxInt32
			}
			num = 10*num + tmp*sign
		} else {
			break
		}
	}
	return num
}

type Offer67 struct{}
type TestCaseOffer67 struct {
	str string
	ans int
}

func (o *Offer67) Run() {
	tests := []TestCaseOffer67{
		{
			str: "42",
			ans: 42,
		},
		{
			str: "  -42",
			ans: -42,
		},
		{
			str: "4193 with words",
			ans: 4193,
		},
		{
			str: "words and 987",
			ans: 0,
		},
		{
			str: "-91283472332",
			ans: -2147483648,
		},
		{
			str: "+-2",
			ans: 0,
		},
		{
			str: "   +0 123",
			ans: 0,
		},
		{
			str: "2147483648",
			ans: 2147483647,
		},
		{
			str: "9223372036854775808",
			ans: 2147483647,
		},
		{
			str: "0-1",
			ans: 0,
		},
		{
			str: "0 123",
			ans: 0,
		},
	}
	for i, t := range tests {
		res := strToInt(t.str)
		if res != t.ans {
			log.Fatalf("[Offer67] test %d failed, expect: %v, got: %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Offer67] All tests passed")
}
