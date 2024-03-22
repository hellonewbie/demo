package four

import "sort"

// 因为这个是有序的，所以直接使用方法来做
func search(nums []int, target int) int {
	lef := sort.SearchInts(nums, target)
	if lef == len(nums) || nums[lef] != target {
		return 0
	}
	rig := sort.SearchInts(nums, target+1) - 1
	return rig - lef + 1
}
