package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"log"
	"time"
)

func main() {
	filename := `./xx.log`
	cofig := tail.Config{
		ReOpen: true, //
		Follow: true,
		//
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	//打开文件读取数据到channel里面
	tails, err := tail.TailFile(filename, cofig)
	if err != nil {
		//%v打印出值的默认格式
		log.Printf("tailfile failed err%v\n", err)
		return
	}
	//开始读取数据
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg", msg.Text)
	}

}
