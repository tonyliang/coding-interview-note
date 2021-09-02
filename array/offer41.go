package array

/*
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例 1：

输入：
["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]
输出：[null,null,null,1.50000,null,2.00000]
示例 2：

输入：
["MedianFinder","addNum","findMedian","addNum","findMedian"]
[[],[2],[],[3],[]]
输出：[null,null,2.00000,null,2.50000]


限制：

最多会对 addNum、findMedian 进行 50000 次调用。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import "github.com/tonyliang/coding-interview-note/helper"

type MedianFinder struct {
	minHeap *helper.Heap
	maxHeap *helper.Heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: helper.NewHeap(), //store the larger half of stream
		maxHeap: helper.NewHeap(), //store the smaller half of stream
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Size() == this.maxHeap.Size() {
		//push to smaller half, pop from smaller half, push to larger half
		this.maxHeap.Push(-num)
		this.minHeap.Push(-this.maxHeap.Pop().(int))
	} else {
		//larger half has more items.
		//push to larger, pop from larger, push to smaller.
		this.minHeap.Push(num)
		this.maxHeap.Push(-this.minHeap.Pop().(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Size() == this.maxHeap.Size() {
		a := this.minHeap.Peek().(int)
		b := -this.maxHeap.Peek().(int)
		return float64(a+b) / float64(2)
	} else {
		return float64(this.minHeap.Peek().(int))
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */