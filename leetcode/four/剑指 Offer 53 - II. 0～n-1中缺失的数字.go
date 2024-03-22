package four

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
			break
		}
	}
	return -1
}
