package main

import (
	"fmt"
	"sync"
)

//使用工人池
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

//缺点：同时扫描过多的主机或端口可能会导致网络或系统限制，造成结果不正确

func main() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
