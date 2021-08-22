package binarysearch

import (
	"fmt"
	"log"
	"math"
)

/*
https://leetcode.com/problems/koko-eating-bananas/
875. Koko Eating Bananas [M]
*/

func minEatingSpeed(piles []int, h int) int {
	min := 1
	max := findMax(piles)
	var left, right, mid, cur int
	left = min
	right = max
	for left < right {
		mid = left + (right-left)/2
		cur = counter(piles, mid)
		if cur > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func counter(arr []int, eat int) int {
	ans := 0
	for _, v := range arr {
		quotient, remainer := v/eat, v%eat
		ans += quotient
		if remainer > 0 {
			ans++
		}
	}
	return ans
}

func findMax(arr []int) int {
	max := math.MinInt64
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

type TestCaseP875 struct {
	id    int
	piles []int
	h     int
	ans   int
}

type P875 struct{}

func (p *P875) Run() {
	test1 := &TestCaseP875{
		id:    1,
		piles: []int{3, 6, 7, 11},
		h:     8,
		ans:   4,
	}
	result1 := minEatingSpeed(test1.piles, test1.h)
	if result1 != test1.ans {
		log.Fatalf("test %d failed, expect: %d, got: %d", test1.id, test1.ans, result1)
	}

	test2 := &TestCaseP875{
		id:    2,
		piles: []int{30, 11, 23, 4, 20},
		h:     5,
		ans:   30,
	}
	result2 := minEatingSpeed(test2.piles, test2.h)
	if result2 != test2.ans {
		log.Fatalf("test %d failed, expect: %d, got: %d", test2.id, test2.ans, result2)
	}

	test3 := &TestCaseP875{
		id:    3,
		piles: []int{30, 11, 23, 4, 20},
		h:     6,
		ans:   23,
	}
	result3 := minEatingSpeed(test3.piles, test3.h)
	if result3 != test3.ans {
		log.Fatalf("test %d failed, expect: %d, got: %d", test3.id, test3.ans, result3)
	}

	test4 := &TestCaseP875{
		id:    4,
		piles: []int{5},
		h:     4,
		ans:   2,
	}
	result4 := minEatingSpeed(test4.piles, test4.h)
	if result4 != test4.ans {
		log.Fatalf("[P875] test %d failed, expect: %d, got: %d", test4.id, test4.ans, result4)
	}

	fmt.Println("[P875] All tests passed.")
}
