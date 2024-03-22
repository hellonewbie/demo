package tailfile

import (
	"08logagent/common"
	"08logagent/kafka"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

type tailTsk struct {
	path  string
	topic string
	tobj  *tail.Tail
}

func newTailTask(path, topic string) *tailTsk {
	tt := &tailTsk{
		path:  path,
		topic: topic,
	}
	return tt
}

func (t *tailTsk) run() {
	logrus.Infof("collect for path :%s is running...", t.path)
	for {
		Line, ok := <-t.tobj.Lines
		if !ok {
			logrus.Warn("tail file close reopen filename:%s\n", t.tobj.Filename)
			time.Sleep(time.Second)
			continue
		}
		//修剪返回字符串的一段，其中删除了剪切集中包含的所有前导和尾随 Unicode 代码点。
		//返回将 s 前后端所有 cutset 包含的 utf-8 码值都去掉的字符串。
		if len(strings.Trim(Line.Text, "/r")) == 0 {
			continue
		}
		//利用通道将同步的代码改为异步
		//把读出来的一行日志包装成kafka里面的msg类型丢到通道中，减少占用内存所以传的是物理地址
		msg := &sarama.ProducerMessage{}
		msg.Topic = t.topic
		msg.Value = sarama.StringEncoder(Line.Text)
		fmt.Println(msg.Topic)
		//丢到通道中
		kafka.MsgChan(msg)
	}
}

func (t *tailTsk) Init() (err error) {
	//allConf里面存了若干个日志的收集项
	//针对每一个日志收集项创建一个对应的tailObj
	config := tail.Config{
		ReOpen: true, //
		Follow: true,
		//
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	//打开文件读取数据到channel里面
	t.tobj, err = tail.TailFile(t.path, config)
	return
}

func Init(allConf []common.CollectEntry) (err error) {
	var wg sync.WaitGroup
	wg.Add(2)
	for _, conf := range allConf {
		tt := newTailTask(conf.Path, conf.Topic) //创建一个日志收集任务
		err := tt.Init()                         //打开文件
		if err != nil {
			logrus.Errorf("create tailobj for path :%s failed,err%v", conf.Path, err)
			continue
		}
		logrus.Infof("create a tail failed task for path %s success", conf.Path)
		//去收集日志吧
		go tt.run()
	}
	wg.Wait()
	return
}
