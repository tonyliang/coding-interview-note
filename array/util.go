package array

func SwapInt(arr []int, i, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}
