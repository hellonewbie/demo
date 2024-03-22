package main

import (
	"fmt"
	"net"
)

//tcp端口扫描

func main() {
	//扫1~1024端口看看有没有开放的
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		//第一个参数要小写表示启动的连接类型，第二个参数表明想要连接的主机，他是单个字符串
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}

}
