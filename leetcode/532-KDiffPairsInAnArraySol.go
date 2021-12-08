package leetcode

func solution532(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n]++
	}
	ans := 0
	for n := range counter {
		if k == 0 && counter[n] > 1 {
			ans++
			continue
		}
		if k > 0 && counter[k+n] > 0 {
			ans++
		}
	}
	return ans
}
