package array

import (
	"fmt"
	"log"
)

/*
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。


示例 1:

输入: [1,2,3,4,5]
输出: True


示例 2:

输入: [0,0,1,2,5]
输出: True


限制：

数组长度为 5

数组的数取值为 [0, 13] .

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isStraight(nums []int) bool {
	quota := make(map[int]int)
	min := 14 //max possible value + 1 given by the question.
	for _, n := range nums {
		quota[n]++
		if min > n && n > 0 {
			min = n
		}
	}
	cur := min
	quota[cur]--
	if quota[cur] == 0 {
		delete(quota, cur)
	}
	for i := 1; i < 5; i++ {
		next := cur + 1
		if _, ok := quota[next]; ok {
			cur = next
			quota[next]--
			if quota[next] == 0 {
				delete(quota, next)
			}
		} else if _, ok := quota[0]; ok {
			cur = next
			quota[0]--
			if quota[0] == 0 {
				delete(quota, 0)
			}
		} else {
			return false
		}
	}
	return true

}

type Offer61 struct{}
type TestCaseOffer61 struct {
	nums []int
	ans  bool
}

func (o *Offer61) Run() {
	tests := []TestCaseOffer61{
		{
			nums: []int{1, 2, 3, 4, 5},
			ans:  true,
		},
		{
			nums: []int{0, 0, 1, 2, 5},
			ans:  true,
		},
		{
			nums: []int{11, 12, 13, 0, 0},
			ans:  true,
		},
		{
			nums: []int{0, 0, 1, 10, 11},
			ans:  false,
		},
	}
	for i, t := range tests {
		res := isStraight(t.nums)
		if res != t.ans {
			log.Fatalf("[Offer61] test %d faild, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer61] All tests passed")
}
