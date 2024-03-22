package main

import (
	"io"
	"log"
	"net"
)

//构建一个端口转发器
func handle(src net.Conn) {
	//destination  目的地址
	dst, err := net.Dial("tcp", "")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)

		}
	}()
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	listerner, err := net.Listen("tcp", "80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	for {
		conn, err := listerner.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
