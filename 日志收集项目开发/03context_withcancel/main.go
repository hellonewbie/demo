package main

import (
	"context"
	"fmt"
	"sync"

	"time"
)

//全局的等待组
var wg sync.WaitGroup

// context.Context是接口类型
//顺手就给大家把接口的概念普及了吧
//Interface是一种类型，和往常语言的接口不一样，它只是用来将对方法进行一个收束。
//然而正是这种收束，使GO语言拥有了基于功能的面向对象。
//定义接口(收束方法)->定义结构体实现方法（new一个结构体实例就面向对象了）->然后这个实例拥有接口的所有功能
func worker(ctx context.Context) {
	defer wg.Done()
	//想要直接退出循环可以加lable
LABLE:
	//一般使用轮询从chan中取值
	for {
		select {
		//func (Context) Done() <-chan struct{} 我们声明通道切片的时候也可以声明结构体类型的 是单向的
		//因为空结构体的宽度是0，占用了0字节的内存空间
		//完成返回一个频道，当代表此上下文完成的工作应被取消时，该频道将关闭
		//Done 通道的关闭可能会异步发生。
		case <-ctx.Done():
			fmt.Println("over")
			break LABLE
		default:
			fmt.Println("worker")
			time.Sleep(time.Second)
		}
	}
	// 如何接收外部命令实现退出
}

func main() {
	//context.WithCancel取消此上下文将释放与其关联的资源
	//上下文：也就是执行任务所需要的相关信息
	//context.Background()返回一个空的Context
	//我们可以用这个空的 Context 作为 goroutine（轻量级线程）的root 节点（看成树的情况下）
	//使用context.WithCancel(parent)函数，创建一个可取消的子Context
	ctx, cancel := context.WithCancel(context.Background())
	//我的理解是这样的这个ctx是父goroutine的一个子goroutine，然后这个cancel绑定了这个子goroutine
	//在要取消上下文释放资源的时候，也就是取消串起来的上下文（ctx充当一个串起来的角色吧）直接调用对应的子cancel函数
	wg.Add(1)
	go worker(ctx)
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 10)
	cancel() //直接调用取消函数释放子goroutine的资源，我感觉就相当于往ctx.Done()里面传了值
	wg.Wait()

}
