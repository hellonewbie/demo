package main

import "strings"

//给定一个字符串数组 words，请计算当两个字符串
//words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。
//假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。

//话不多多说直接暴力开整

func maxProduct(words []string) int {
	var Maxlenth int = 0
	var Word1 string
	for i := range words {
		Word1 = words[i]
		for j := i + 1; j < len(words); j++ {
			Word2 := words[j]
			if !haveSameChar(Word1, Word2) {
				if len(Word1)*len(Word2) > Maxlenth {
					Maxlenth = len(Word1) * len(Word2)
				}
			}
		}
	}
	return Maxlenth
}

func haveSameChar(Word1 string, Word2 string) bool {
	for _, v := range Word1 {
		//使用strings包下的index如果word2中存在该字符会返回该字符的下标
		//时间复杂度为O()
		if strings.Index(Word2, string(v)) >= 0 {
			return true
		}
	}
	return false
}

//降低时间复杂度
func hasSameChar2(word1 string, word2 string) bool {
	var count = [26]int{}
	for _, c := range word1 {
		count[c-'a'] = 1
	}
	for _, c := range word2 {
		if count[c-'a'] == 1 {
			return true
		}
	}
	return false
}
