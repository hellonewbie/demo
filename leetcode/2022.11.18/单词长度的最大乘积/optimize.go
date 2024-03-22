package main

//给定一个字符串数组 words，请计算当两个字符串
//words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。
//假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。

// 位运算 + 预计算
// 时间复杂度：O((m + n)* n)
// 空间复杂度：O(n)
func maxProduct1(words []string) int {
	// O(mn)
	var ant int = 0
	n := len(words)
	masks := make([]int, 26)
	for i := 0; i < n; i++ {
		bitMask := 0
		for _, c := range words[i] {
			bitMask |= 1 << (c - 'a')
		}
		masks[i] = bitMask
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if masks[i]&masks[j] == 0 && len(words[i])*len(words[j]) > ant {
				ant = len(words[i]) * len(words[j])
			}
		}
	}
	return ant
}
