package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	//生产者配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll          //ACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner //分区
	//2封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("666666")
	//3、连接kafka,放的是string类型的切片所以放多个连接多个kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.110.131:9092"}, config)
	if err != nil {
		fmt.Println("produer closed err", err)
		return
	}
	defer client.Close()
	//4、发消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed ", err)
		return
	}
	fmt.Printf("pid%v offset:%v\n", pid, offset)
}
