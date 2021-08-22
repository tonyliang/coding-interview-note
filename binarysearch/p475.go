package binarysearch

import (
	"fmt"
	"log"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	res := 0

	for _, h := range houses {
		closest := binarySearchClosestHeater(heaters, h)
		res = Max(res, Abs(h-closest))
	}

	return res
}

// to find the right spot for target in heaters
func binarySearchClosestHeater(heaters []int, target int) int {
	left := 0
	right := len(heaters) - 1
	for left+1 < right {
		mid := left + (right-left)/2
		if heaters[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	// NOTE: exit criterion is left + 1 == right
	if target <= heaters[left] {
		return heaters[left]
	}
	if target >= heaters[right] {
		return heaters[right]
	}
	if Abs(heaters[left]-target) > Abs(heaters[right]-target) {
		return heaters[right]
	} else {
		return heaters[left]
	}
}

// func findRadius(houses []int, heaters []int) int {
// 	min := 0 //Min(houses[0], heaters[0])
// 	sort.Ints(heaters)
// 	max := Max(houses[len(houses)-1], heaters[len(heaters)-1])
// 	var left, right int
// 	left = min
// 	right = max

// 	for left < right {
// 		mid := left + (right-left)/2
// 		spans := makeSpans(heaters, mid)
// 		isCovered := checkCover(houses, spans)
// 		if isCovered {
// 			right = mid
// 		} else {
// 			left = mid + 1
// 		}
// 	}
// 	return left
// }

// type Span struct {
// 	start int
// 	end   int
// }

// func makeSpans(heaters []int, radius int) []Span {
// 	result := make([]Span, 0)
// 	for _, h := range heaters {
// 		s := Span{
// 			start: Max(h-radius, 0),
// 			end:   h + radius,
// 		}
// 		result = append(result, s)
// 	}
// 	return result
// }

// func checkCover(houses []int, spans []Span) bool {
// 	for _, h := range houses {
// 		covered := false
// 		for _, s := range spans {
// 			if s.start <= h && s.end >= h {
// 				covered = true
// 				break
// 			}
// 		}
// 		if !covered {
// 			return false
// 		}
// 	}
// 	return true
// }

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type TestCaseP475 struct {
	houses  []int
	heaters []int
	ans     int
}

type P475 struct{}

func (p *P475) Run() {
	tests := []TestCaseP475{
		// {
		// 	houses:  []int{1, 2, 3},
		// 	heaters: []int{2},
		// 	ans:     1,
		// },
		{
			houses:  []int{1, 2, 3, 4},
			heaters: []int{1, 4},
			ans:     1,
		},
		{
			houses:  []int{1, 2, 3},
			heaters: []int{2},
			ans:     1,
		},
		{
			houses:  []int{1, 5},
			heaters: []int{2},
			ans:     3,
		},
		{
			houses:  []int{1, 2, 3, 9},
			heaters: []int{5, 8},
			ans:     4,
		},
		{
			houses:  []int{1},
			heaters: []int{1},
			ans:     0,
		},
		{
			houses:  []int{1},
			heaters: []int{1, 2, 3, 4},
			ans:     0,
		},
	}
	var result int
	for i, t := range tests {
		result = findRadius(t.houses, t.heaters)
		if result != t.ans {
			log.Fatalf("[P475] Test %d failed. Expect: %d, Got: %d", i, t.ans, result)
		}
	}
	fmt.Println("[P475] All tests passed")
}
