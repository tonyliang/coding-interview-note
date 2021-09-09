package helper

import (
	"fmt"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	var tests = []struct {
		op       string
		val      int
		wantHead int
		wantTail int
	}{
		{"ADD", 1, 1, 1},
		{"SIZE", 1, -1, -1},
		{"ADD", 2, 2, 1},
		{"SIZE", 2, -1, -1},
		{"ADD", 3, 3, 1},
		{"SIZE", 3, -1, -1},
		{"ADD", 4, 4, 2},
		{"SIZE", 3, -1, -1},
		{"REMOVE", 2, 4, 3},
		{"SIZE", 2, -1, -1},
		{"REMOVE", 3, 4, 4},
		{"SIZE", 1, -1, -1},
		{"REMOVE", 4, -1, -1},
		{"SIZE", 0, -1, -1},
		//part 2
		{"ADD", 1, 1, 1},
		{"ADD", 2, 2, 1},
		{"ADD", 3, 3, 1},
		{"ADD", 4, 4, 2},
		{"ADD", 5, 5, 3},
		{"ADD", 6, 6, 4},
		{"SIZE", 3, -1, -1},
		{"REMOVE", 4, 6, 5},
		{"REMOVE", 5, 6, 6},
		{"ADD", 5, 5, 6},
		{"ADD", 4, 4, 6},
		{"ADD", 3, 3, 5},
	}
	rb := NewRingBuffer(3)
	for i, tt := range tests {
		testname := fmt.Sprintf("%d,%v", i, tt.op)
		t.Run(testname, func(t *testing.T) {
			switch tt.op {
			case "ADD":
				rb.Add(tt.val)
				head := rb.GetHead()
				tail := rb.GetTail()
				if head.(int) != tt.wantHead {
					t.Errorf("got %d, want %d", head, tt.wantHead)
				}
				if tail.(int) != tt.wantTail {
					t.Errorf("got %d, want %d", tail, tt.wantTail)
				}
				fmt.Printf("[pass] %d, %v\n", i, tt)
			case "SIZE":
				size := rb.Count()
				if size != tt.val {
					t.Errorf("got %d, want %d", size, tt.val)
				}
				fmt.Printf("[pass] %d, %v\n", i, tt)
			case "REMOVE":
				removed := rb.Remove()
				if removed.(int) != tt.val {
					t.Errorf("got %d, want %d", removed, tt.val)
				}
				head := rb.GetHead()
				tail := rb.GetTail()
				if head == nil && tt.wantHead != -1 {
					t.Errorf("got %d, want %d", head, tt.wantHead)
				}
				if tail == nil && tt.wantTail != -1 {
					t.Errorf("got %d, want %d", tail, tt.wantTail)
				}
				if head != nil && head.(int) != tt.wantHead {
					t.Errorf("got %d, want %d", head, tt.wantHead)
				}
				if tail != nil && tail.(int) != tt.wantTail {
					t.Errorf("got %d, want %d", tail, tt.wantTail)
				}
				fmt.Printf("[pass] %d, %v\n", i, tt)
			}
		})
	}
}
