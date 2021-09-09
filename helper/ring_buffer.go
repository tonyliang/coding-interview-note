package helper

import "log"

type RingBuffer struct {
	buffer []interface{}
	head   int
	tail   int
	size   int
	count  int
}

/*
The idea is that head is alway pointing at the next availabe slot for insertion.
Therefore, the latest item added is at head-1,
The oldest item added is at tail,
When head == tail, the state can be either empty or full. If the buffer is full and we need to advance tail by one.
We also use count and size to verify if the buffer is empty state or full.
*/

func NewRingBuffer(size int) *RingBuffer {
	if size < 2 {
		log.Fatal("Size of RingBuffer cannot be less than 2")
	}
	buffer := make([]interface{}, size)
	return &RingBuffer{
		buffer: buffer,
		head:   0,
		tail:   0,
		size:   size,
		count:  0,
	}
}

func (rb *RingBuffer) Count() int {
	return rb.count
}

func (rb *RingBuffer) IsFull() bool {
	return rb.count == rb.size
}

func (rb *RingBuffer) GetHead() interface{} {
	if rb.count == 0 {
		return nil
	}
	pos := ((rb.head + rb.size) - 1) % rb.size
	return rb.buffer[pos]
}

func (rb *RingBuffer) GetTail() interface{} {
	if rb.count == 0 {
		return nil
	}
	return rb.buffer[rb.tail]
}

func (rb *RingBuffer) Add(item interface{}) {
	if rb.IsFull() {
		rb.tail = (rb.tail + 1) % rb.size
	} else {
		rb.count++
	}
	rb.buffer[rb.head] = item
	rb.head = (rb.head + 1) % rb.size
}

func (rb *RingBuffer) Remove() interface{} {
	if rb.count == 0 {
		return nil
	}
	removed := rb.buffer[rb.tail]
	rb.tail = (rb.tail + 1) % rb.size
	rb.count--
	return removed
}
