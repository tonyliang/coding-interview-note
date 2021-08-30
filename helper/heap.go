package helper

import (
	"fmt"
	"log"
)

type Heap struct {
	data []interface{}
}

func (h *Heap) Peek() interface{} {
	if len(h.data) == 0 {
		log.Fatal("Trying to access empty list")
	}
	return h.data[0]
}

func (h *Heap) Push(el interface{}) {
	h.data = append(h.data, el)
	h.bubbleUp()
}

func (h *Heap) bubbleUp() {
	cur := len(h.data) - 1
	parent := h.getParentIdx(cur)
	for cur > 0 && h.less(h.data[cur], h.data[parent]) {
		h.swap(cur, parent)
		cur = parent
		parent = h.getParentIdx(cur)
	}
}

func (h *Heap) Pop() interface{} {
	if len(h.data) == 0 {
		return nil
	}
	toReturn := h.data[0]
	h.swap(0, len(h.data)-1)
	h.data = h.data[0 : len(h.data)-1]
	h.sink(0)
	return toReturn
}

func (h *Heap) sink(cur int) {
	if cur >= len(h.data) {
		return
	}
	curV := h.data[cur]
	left, leftV := h.getLeftChild(cur)
	right, rightV := h.getRightChild(cur)
	if leftV != nil && rightV != nil {
		if h.less(leftV, rightV) {
			if h.less(leftV, curV) {
				h.swap(cur, left)
				h.sink(left)
			}
		} else {
			if h.less(rightV, curV) {
				h.swap(cur, right)
				h.sink(right)
			}
		}
	} else if leftV != nil {
		if h.less(leftV, curV) {
			h.swap(cur, left)
			h.sink(left)
		}
	} else if rightV != nil {
		if h.less(rightV, curV) {
			h.swap(cur, right)
			h.sink(right)
		}
	} else {
		return
	}
}

func (h *Heap) swap(a, b int) {
	h.data[a], h.data[b] = h.data[b], h.data[a]
}

func (h *Heap) less(a, b interface{}) bool {
	switch a.(type) {
	case int:
		return int(a.(int)) < int(b.(int))
	case float64:
		return float64(a.(float64)) < float64(b.(float64))
	default:
		return string(a.(string)) < string(b.(string))
	}
}

func (h *Heap) getParentIdx(idx int) int {
	return (idx - 1) / 2
}

func (h *Heap) getLeftChild(idx int) (int, interface{}) {
	i := (idx * 2) + 1
	if i < len(h.data) {
		return i, h.data[i]
	}
	return -1, nil
}

func (h *Heap) getRightChild(idx int) (int, interface{}) {
	i := (idx * 2) + 2
	if i < len(h.data) {
		return i, h.data[i]
	}
	return -1, nil
}

func (h *Heap) Size() int {
	return len(h.data)
}

func NewHeap() *Heap {
	return &Heap{
		data: make([]interface{}, 0),
	}
}

func TestHeap() {
	h := NewHeap()

	data1 := []int{30, 10, 20, 80}
	for _, v := range data1 {
		h.Push(v)
	}
	res1 := []int{10, 20, 30, 80}
	for i := 0; i < len(data1); i++ {
		res := h.Pop().(int)
		if res != res1[i] {
			log.Fatalf("test 1 failed, expect index %d to be %d, got %d\n", i, res1[i], res)
		}
	}

	data2 := []int{1, 2, 3, 5, 9, 0}
	for _, v := range data2 {
		h.Push(v)
	}
	res2 := []int{0, 1, 2, 3, 5, 9}
	for i := 0; i < len(data2); i++ {
		res := h.Pop().(int)
		if res != res2[i] {
			log.Fatalf("test 2 failed, expect index %d to be %d, got %d\n", i, res2[i], res)
		}
	}

	data3 := []int{3, 3, 2, 2, 9, 9, 9, 1}
	for _, v := range data3 {
		h.Push(v)
	}
	res3 := []int{1, 2, 2, 3, 3, 9, 9, 9}
	for i := 0; i < len(data3); i++ {
		res := h.Pop().(int)
		if res != res3[i] {
			log.Fatalf("test 3 failed, expect index %d to be %d, got %d\n", i, res3[i], res)
		}
	}
	r := -1
	h.Push(100)
	h.Push(200)
	h.Push(300)
	r = h.Pop().(int)
	if r != 100 {
		log.Fatal("failed, should be 100")
	}
	h.Push(10)
	r = h.Pop().(int)
	if r != 10 {
		log.Fatal("failed, should be 10")
	}
	h.Push(-1)
	h.Push(-2)
	r = h.Pop().(int)
	if r != -2 {
		log.Fatalf("failed, should be -2, got %v", r)
	}
	r = h.Pop().(int)
	if r != -1 {
		log.Fatal("failed, should be -2")
	}
	r = h.Pop().(int)
	if r != 200 {
		log.Fatal("failed, shold be 200")
	}
	r = h.Pop().(int)
	if r != 300 {
		log.Fatal("failed, should be 300")
	}

	fmt.Println("All tests passed")
}
