package stack

import (
	"fmt"
	"log"
)

type MinStack struct {
	in  []int //in stores all elements
	out []int //out stores min element in decreasing order
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.in = append(this.in, x)
	if len(this.out) == 0 || this.out[len(this.out)-1] >= x {
		this.out = append(this.out, x)
	}
}

func (this *MinStack) Pop() {
	r := this.in[len(this.in)-1]
	this.in = this.in[:len(this.in)-1]
	if r == this.out[len(this.out)-1] {
		this.out = this.out[:len(this.out)-1]
	}
}

func (this *MinStack) Top() int {
	return this.in[len(this.in)-1]
}

func (this *MinStack) Min() int {
	return this.out[len(this.out)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

type P155 struct{}
type TestCaseP155 struct {
	op  string
	val interface{}
	ans interface{}
}

func (p *P155) Run() {
	tests := []TestCaseP155{
		{
			op:  "push",
			val: -2,
			ans: nil,
		},
		{
			op:  "push",
			val: 0,
			ans: nil,
		},
		{
			op:  "push",
			val: -3,
			ans: nil,
		},
		{
			op:  "min",
			val: nil,
			ans: -3,
		},
		{
			op:  "pop",
			val: nil,
			ans: nil,
		},
		{
			op:  "top",
			val: nil,
			ans: 0,
		},
		{
			op:  "min",
			val: nil,
			ans: -2,
		},
	}
	obj := Constructor()
	var res interface{}
	for i, t := range tests {
		switch op := t.op; op {
		case "push":
			obj.Push(t.val.(int))
			res = nil
		case "pop":
			obj.Pop()
			res = nil
		case "min":
			res = obj.Min()
		case "top":
			res = obj.Top()
		}
		if res != t.ans {
			log.Fatalf("[P155] Test %d failed. Expect %v, Got %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[P155] All tests passed")
}
