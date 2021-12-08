package leetcode

import (
	"sort"

	"github.com/tonyliang/coding-interview-note/helper"
)

type SortableEvents []Event218

type Event218 struct {
	X     int
	H     int
	Begin bool
	End   bool
}

func (e *Event218) isBegin() bool {
	return e.Begin
}

func (e *Event218) isEnd() bool {
	return e.End
}

func (s SortableEvents) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableEvents) Less(i, j int) bool {
	if s[i].X == s[j].X {
		if s[i].isBegin() && s[j].isEnd() {
			return true
		}
		if s[j].isBegin() && s[i].isEnd() {
			return false
		}
		if s[i].isBegin() && s[j].isBegin() {
			return s[j].H < s[i].H
		}
		if s[i].isEnd() && s[j].isEnd() {
			return s[i].H < s[j].H
		}
	}
	return s[i].X < s[j].X
}

func (s SortableEvents) Len() int {
	return len(s)
}

func solution218(buildings [][]int) [][]int {
	events := make(SortableEvents, 0)
	for _, b := range buildings {
		events = append(events, Event218{X: b[0], H: b[2], Begin: true})
		events = append(events, Event218{X: b[1], H: b[2], End: true})
	}
	sort.Sort(events)
	var root *helper.BSTNode
	ans := make([][]int, 0)
	for _, e := range events {
		if e.isBegin() {
			maxNode := helper.BSTSearchMax(root)
			if maxNode == nil || e.H > maxNode.Key {
				ans = append(ans, []int{e.X, e.H})
			}
			root = helper.BSTInsert(root, e.H)
		}
		if e.isEnd() {
			root = helper.BSTDelete(root, e.H)
			maxNode := helper.BSTSearchMax(root)
			if maxNode == nil || maxNode.Key < e.H {
				if maxNode == nil {
					ans = append(ans, []int{e.X, 0})
				} else {
					ans = append(ans, []int{e.X, maxNode.Key})
				}
			}
		}

	}
	return ans
}
