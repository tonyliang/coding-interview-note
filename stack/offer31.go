package stack

import (
	"fmt"
	"log"
)

/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。



示例 1：

输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
示例 2：

输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	i, j := 0, 0
	for i < len(pushed) {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[0 : len(stack)-1]
			j++
		}
		i++
	}
	return len(stack) == 0
}

type Offer31 struct{}
type TestCaseOffer31 struct {
	pushed []int
	popped []int
	ans    bool
}

func (o *Offer31) Run() {
	tests := []TestCaseOffer31{
		{
			pushed: []int{1, 2, 3, 4, 5},
			popped: []int{4, 5, 3, 2, 1},
			ans:    true,
		},
		{
			pushed: []int{1, 2, 3, 4, 5},
			popped: []int{2, 5, 4, 3, 1},
			ans:    true,
		},
		{
			pushed: []int{1, 2, 3, 4, 5},
			popped: []int{4, 3, 5, 1, 2},
			ans:    false,
		},
		{
			pushed: []int{1, 0},
			popped: []int{1, 0},
			ans:    true,
		},
		{
			pushed: []int{0, 2, 1},
			popped: []int{0, 1, 2},
			ans:    true,
		},
		{
			pushed: []int{2, 1, 0},
			popped: []int{1, 2, 0},
			ans:    true,
		},
	}
	for i, t := range tests {
		res := validateStackSequences(t.pushed, t.popped)
		if res != t.ans {
			log.Fatalf("[Offer31] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer31] All tests passed")
}
