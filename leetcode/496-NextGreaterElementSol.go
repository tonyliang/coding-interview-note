package leetcode

func solution496(nums1, nums2 []int) []int {
	nextG := make(map[int]int)
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			nextG[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		if n, ok := nextG[v]; ok {
			res[i] = n
		} else {
			res[i] = -1
		}
	}
	return res
}

func solution496_2(nums1, nums2 []int) []int {
	indexes := make(map[int]int)
	for i, v := range nums2 {
		indexes[v] = i
	}
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		j, ok := indexes[v]
		if !ok {
			res[i] = -1
			continue
		}

		found := false
		for k := j + 1; k < len(nums2); k++ {
			if nums2[k] > v {
				res[i] = nums2[k]
				found = true
				break
			}
		}
		if !found {
			res[i] = -1
		}
	}
	return res
}
