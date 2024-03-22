package five

//问题出在high的移动问题上，,high只能移到mid的位置，如果移到mid-1可能会越界
//但是low不会出现这种情况，因为/保留的整数位所以右边一定还有一位是可以让low移动过去的
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for high > low {
		mid := (high + low) / 2
		if numbers[mid] > numbers[high] {
			low = mid + 1
		} else if numbers[mid] < numbers[high] {
			high = mid
		} else {
			high = high - 1
		}
	}
	return numbers[low]
}
