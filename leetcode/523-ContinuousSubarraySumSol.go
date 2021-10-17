package leetcode

func solution523(nums []int, k int) bool {
	seen := make(map[int]int)
	seen[0] = -1 //in case the sum of the entire array is a multiple of k
	sum := 0
	for i, n := range nums {
		sum += n
		if v, ok := seen[sum%k]; ok {
			if i-2 >= v {
				return true
			}
		} else {
			seen[sum%k] = i
		}
	}
	return false
}
