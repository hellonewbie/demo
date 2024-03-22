package main

//func countBits(n int) []int {
//	var Ret []int
//
//	for i := 0; i <= n; i++ {
//		StrI := strconv.FormatInt(int64(i), 2)
//		var num int = 0
//		for j := 0; j < len(StrI); j++ {
//
//			if int(StrI[len(StrI)-j-1]-'0') == 1 {
//				num++
//			}
//		}
//		Ret = append(Ret, num)
//	}
//	return Ret
//}
func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}
func main() {

}
