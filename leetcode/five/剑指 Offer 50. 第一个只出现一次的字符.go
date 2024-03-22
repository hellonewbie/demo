package five

import (
	"fmt"
	"reflect"
	"unsafe"
)

//算法巨垃圾然后性能巨差，有待改进
//然后忘了遍历map是无序的，这个时候我们可以借助一个辅助的切片保存顺序来返回第一个只有一个的字符
//判断map是否为空可以直接使用map的性质 value,ok:=map[]
func firstUniqChar(s string) byte {
	hashmap := make(map[byte]bool, 0)
	xuliehua := make([]byte, 0)
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*[]byte)(unsafe.Pointer(sh))
	for _, v := range *bh {
		fmt.Println(string(v))
		xuliehua = append(xuliehua, v)
		_, ok := hashmap[v]
		if !ok {
			hashmap[v] = true
		} else {
			hashmap[v] = false
		}
	}
	for _, v := range xuliehua {
		if hashmap[v] == true {
			return v
		}
	}
	return ' '
}

func main() {
	firstUniqChar("leetcode")
}
