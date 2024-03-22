package main

import (
	"fmt"
	"sync"

	"time"
)

//全局的等待组
var wg sync.WaitGroup

// 初始的例子

func worker(ch <-chan bool) {
	defer wg.Done()
	//想要直接退出循环可以加lable
LABLE:
	//一般使用轮询从chan中取值
	for {
		select {
		case <-ch:
			break LABLE
		default:
			fmt.Println("worker")
			time.Sleep(time.Second)
		}
	}
	// 如何接收外部命令实现退出
}

func main() {
	//创建一个通道变量 后面跟着的是在通道里面传输的值的类型
	//var exitChan chan  bool
	//make和new的区别:
	//new函数返回值返回的是指针（多用于基本数据类型初始化bool、string、int）
	//make的返回值类型（多用于slice、map、channel） 后面参数0 就是带不带缓冲区的区别
	var exitChan = make(chan bool, 1)
	wg.Add(1)
	go worker(exitChan)
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 10)
	exitChan <- true
	wg.Wait()
	fmt.Println("over")
}
