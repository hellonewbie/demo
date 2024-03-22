package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var (
	Client  sarama.SyncProducer
	msgChan chan *sarama.ProducerMessage
)

func Init(address []string, chanSize int64) (err error) {
	//生产者配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll          //ACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner //分区
	//3、连接kafka,放的是string类型的切片所以放多个连接多个kafka
	Client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		logrus.Error("produer closed err", err)
		return
	}
	msgChan = make(chan *sarama.ProducerMessage, chanSize)
	//起一个后台的goroutine从msgChan中读数据
	go sendMsg()

	return
}

//从MsgChan 中读取msg，发送给kafka
func sendMsg() {
	for {
		select {
		case msg := <-msgChan:
			pid, offset, err := Client.SendMessage(msg)
			if err != nil {
				logrus.Warningf("send msg failed :%v", err)
				return
			}
			logrus.Infof("send message to kafka success,pid:%v offset:%v", pid, offset)
		}
	}

}

//定义一个函数向外暴露msgChan
func MsgChan(msg *sarama.ProducerMessage) {
	msgChan <- msg
}
