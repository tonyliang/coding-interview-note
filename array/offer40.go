package array

import (
	"fmt"
	"log"
	"sort"

	"github.com/tonyliang/coding-interview-note/helper"
)

/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。



示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]

limit:
- 0 <= k <= arr.length <= 10000
- 0 <= arr[i] <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
solution 1: MaxHeap, time complexity: O(N*LogK)
solution 2: Quick Select, when we located the Kth smallest, the K-1 items including the Kth are the answer.
Time complexity: O(N)
solution 3: Counting sort, because the size of arr[i] is fixed from 0 to 10000. Time complexity: O(N)
*/

func getLeastNumbersSolution1(arr []int, k int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	if len(arr) == k {
		return arr
	}
	heap := helper.NewHeap()
	for _, v := range arr {
		heap.Push(-v)
		if heap.Size() > k {
			heap.Pop()
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = -heap.Pop().(int)
	}
	return res
}

func getLeastNumbersSolution3(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	if len(arr) == k {
		return arr
	}
	counter := make([]int, 10001)
	for _, v := range arr {
		counter[v]++
	}
	res := make([]int, k)

	for i := 0; i <= 10000; i++ {
		for counter[i] > 0 {
			res[k-1] = i
			counter[i]--
			k--
			if k == 0 {
				return res
			}
		}
	}
	return nil
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	quickSelect(arr, 0, len(arr)-1, k-1)
	return arr[0:k]
}

func quickSelect(arr []int, start, end, k int) {
	if start < end {
		pivot := partition(arr, start, end)
		if pivot == k {
			return
		}
		if pivot > k {
			quickSelect(arr, start, pivot-1, k)
		} else {
			quickSelect(arr, pivot+1, end, k)
		}
	}
}

func partition(arr []int, start, end int) int {
	pivotVal := arr[end]
	partitionIdx := start
	for i := start; i < end; i++ {
		if arr[i] < pivotVal {
			arr[i], arr[partitionIdx] = arr[partitionIdx], arr[i]
			partitionIdx++
		}
	}
	arr[end], arr[partitionIdx] = arr[partitionIdx], arr[end]
	return partitionIdx
}

type Offer40 struct{}
type TestCaseOffer40 struct {
	arr []int
	k   int
	ans []int
}

func (o *Offer40) Run() {
	tests := []TestCaseOffer40{
		{
			arr: []int{3, 2, 1},
			k:   2,
			ans: []int{1, 2},
		},
		{
			arr: []int{0, 1, 2, 1},
			k:   1,
			ans: []int{0},
		},
		{
			arr: []int{3, 2, 1, 5},
			k:   0,
			ans: []int{},
		},
		{
			arr: []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4},
			k:   8,
			ans: []int{0, 0, 1, 1, 2, 2, 2, 3},
		},
	}
	for i, t := range tests {
		res := getLeastNumbers(t.arr, t.k)
		sort.Ints(res)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Offer40] test %d failed, expect %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer40] All tests passed")
}
