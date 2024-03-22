package main

import (
	"fmt"
	"net"
	"sync"
)

//缺点：同时扫描过多的主机或端口可能会导致网络或系统限制，造成结果不正确

func main() {
	var wg sync.WaitGroup
	//扫1~1024端口看看有没有开放的
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			//想要某一步最后执行，一般可以使用延迟调用
			defer wg.Done()
			address := fmt.Sprintf("127.0.0.1:%d", j)
			//第一个参数要小写表示启动的连接类型，第二个参数表明想要连接的主机，他是单个字符串
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
		wg.Wait()
	}
	fmt.Printf("scanner port finish")
}
