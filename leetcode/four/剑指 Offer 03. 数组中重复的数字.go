package four

//使用hashmap哈希表来实现元素的查找
func findRepeatNumber(nums []int) int {
	find := make(map[int]bool)
	for _, v := range nums {
		if find[v] != true {
			find[v] = true
		} else if find[v] == true {
			return v
			break
		}
	}
	return -1
}
