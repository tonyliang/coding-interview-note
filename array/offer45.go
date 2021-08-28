package array

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

/*
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。



示例 1:

输入: [10,2]
输出: "102"
示例 2:

输入: [3,30,34,5,9]
输出: "3033459"


提示:

0 < nums.length <= 100
说明:

输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type MySort []string

func (m MySort) Len() int {
	return len(m)
}

func (m MySort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MySort) Less(i, j int) bool {
	if m[i] == "0" {
		return true
	}
	if m[j] == "0" {
		return false
	}
	if m[i]+m[j] < m[j]+m[i] {
		return true
	}
	return false
}

func minNumber(nums []int) string {
	if nums == nil {
		return ""
	}
	N := len(nums)
	array := make([]string, N)
	for i, v := range nums {
		array[i] = strconv.Itoa(v)
	}
	sorted := MySort(array)
	sort.Sort(sorted)
	return strings.Join(sorted, "")
}

/* simplified version

func minNumber(nums []int) string {
    // 实际上是排序题 即字符串形式 x+y<y+x表示x<y 否则就是x>y
    sort.Slice(nums,func(i,j int)bool{
        return strconv.Itoa(nums[i])+strconv.Itoa(nums[j])<strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
    })
    ret:=""
    for _,v:= range nums{
        ret+=strconv.Itoa(v)
    }
    return ret
}

*/
type Offer45 struct{}
type TestCaseOffer45 struct {
	nums []int
	ans  string
}

func (o *Offer45) Run() {
	tests := []TestCaseOffer45{
		{
			nums: []int{10, 2},
			ans:  "102",
		},
		{
			nums: []int{3, 30, 34, 5, 9},
			ans:  "3033459",
		},
		{
			nums: []int{3, 300, 30, 0, 0},
			ans:  "00300303",
		},
	}
	for i, t := range tests {
		res := minNumber(t.nums)
		if res != t.ans {
			log.Fatalf("[Offer45] test %d failed, expect: %v, got: %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Offer45] All tests passed")
}
