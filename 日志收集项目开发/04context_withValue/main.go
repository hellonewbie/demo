package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type TraceCode string //为了所提供的键是可以比较的，所以我们声明一个自己的类型

var wg sync.WaitGroup

func work(ctx context.Context) {
	KEY := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(KEY).(string) //把拿出来的值进行断言只有断对了类型才有值返回
	fmt.Println(ctx.Value(KEY))
	if !ok {
		log.Print("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "123456")
	wg.Add(1)
	go work(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
