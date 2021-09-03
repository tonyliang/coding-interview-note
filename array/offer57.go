package array

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。



示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]


限制：

1 <= target <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//Time: O(N), Space: O(N)
func findContinuousSequenceX(target int) [][]int {
	data := make([]int, target)
	for i := 1; i <= target; i++ {
		data[i-1] = i
	}
	left, right, total := 0, 0, 0
	ans := make([][]int, 0)

	for right < len(data) {
		total += data[right]
		for total > target {
			total -= data[left]
			left++
		}
		if total == target && left < right {
			ans = append(ans, data[left:right+1])
		}
		right++
	}
	return ans
}

//Time: O(N), Space: O(1)
func findContinuousSequence(target int) [][]int {
	if target <= 2 {
		return nil
	}
	ans := make([][]int, 0)
	left, right := 1, 2
	for left <= right {
		sum := (right + left) * (right - left + 1) / 2
		if sum > target {
			left++
		} else if sum < target {
			right++
		} else {
			//sum == target
			if left < right {
				ans = append(ans, enlistHelper(left, right))
			}
			//there won't be 2nd answer starts at left again, so left++
			left++
		}
	}
	return ans
}

func enlistHelper(a, b int) []int {
	r := make([]int, b-a+1)
	cur := a
	for i := 0; i < len(r); i++ {
		r[i] = cur
		cur++
	}
	return r
}

type Offer57 struct{}
type TestCaseOffer57 struct {
	target int
	ans    [][]int
}

func (o *Offer57) Run() {
	tests := []TestCaseOffer57{
		{
			target: 9,
			ans:    [][]int{{2, 3, 4}, {4, 5}},
		},
		{
			target: 15,
			ans:    [][]int{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}},
		},
	}
	for i, t := range tests {
		res := findContinuousSequence(t.target)
		if len(res) != len(t.ans) {
			log.Fatalf("[Offer57] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
		for j, v := range res {
			if !helper.Equal(v, t.ans[j]) {
				log.Fatalf("[Offer57] test %d failed, expect: %v, got: %v", i, t.ans[j], v)
			}
		}
	}
	fmt.Println("[Offer57] All tests passed")
}
