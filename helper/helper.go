package helper

import "fmt"

const intSize = 32

const (
	MaxInt    = 1<<(intSize-1) - 1
	MinInt    = -1 << (intSize - 1)
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint   = 1<<intSize - 1
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

func Equal(a, b interface{}) bool {
	switch a.(type) {
	case [][]int:
		aa := a.([][]int)
		bb := b.([][]int)
		if len(aa) != len(bb) {
			return false
		}
		for i := 0; i < len(aa); i++ {
			for j := 0; j < len(aa[0]); j++ {
				if aa[i][j] != bb[i][j] {
					return false
				}
			}
		}
		return true
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
