package helper

import "fmt"

func Equal(a, b interface{}) bool {
	switch a.(type) {
	case []int:
		aa := a.([]int)
		bb := b.([]int)
		if len(aa) != len(bb) {
			return false
		}
		for i := 0; i < len(aa); i++ {
			if aa[i] != bb[i] {
				return false
			}
		}
		return true
	case []string:
		aa := a.([]string)
		bb := b.([]string)
		if len(aa) != len(bb) {
			return false
		}
		for i := 0; i < len(aa); i++ {
			if aa[i] != bb[i] {
				return false
			}
		}
		return true
	default:
		fmt.Println("nope")
	}
	return false
}
