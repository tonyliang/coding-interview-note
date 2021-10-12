package leetcode

/*
Monotonic stack principles:
1. stack should look like 1,2,4,6,11,..
   or  11,6,4,2,1...
2. to add a new item on to stack, pop until the monotonic order is maintained.
3. process answer while popping
4. store index in the stack and fetch real value by mapping (index to value via original input array)

*/
func solution503(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}
	res := make([]int, len(input))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	stack := make([]int, 0)
	for i := 0; i < len(input)*2; i++ {
		index := i % len(input)
		num := input[index]
		for len(stack) > 0 && input[stack[len(stack)-1]] < num {
			poppedIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[poppedIdx] = num
		}
		stack = append(stack, index)
	}
	return res
}
