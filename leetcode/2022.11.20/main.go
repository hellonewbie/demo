package main

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		dif := target - numbers[i]
		if _, ok := m[dif]; ok {
			return []int{m[dif], i}
		}
		m[numbers[i]] = i
	}
	return []int{-1, -1}
}
