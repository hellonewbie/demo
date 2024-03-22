package main

import (
	"fmt"
	"unsafe"
)

//剑指 Offer II 004. 只出现一次的数字
//时间复杂度O(n)

//func singleNumber(nums []int) int {
//	var num int
//	var exam map[int]bool
//	for _, v := range nums {
//		if _, ok := exam[v]; !ok {
//			exam[v] = true
//		} else {
//			exam[v] = false
//		}
//	}
//	for _, v := range nums {
//		if exam[v] {
//			num = v
//		}
//	}
//	return num
//}

//
func main() {
	var exam map[int]bool

	exam1 := make(map[int]bool)
	fmt.Println(unsafe.Sizeof(exam), unsafe.Sizeof(exam1))
}
