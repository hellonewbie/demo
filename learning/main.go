package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func MD5(s string) string {
	c := md5.New()
	c.Write([]byte(s))
	cipherStr := c.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//result, _, err := reader.ReadLine()
	//if err != nil {
	//	fmt.Println("获取失败")
	//}
	//arr := strings.Split(string(result), " ")
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}
	fmt.Println(string(MD5("123456")))
}
