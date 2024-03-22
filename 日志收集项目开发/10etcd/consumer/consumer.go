package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	consumer, err := sarama.NewConsumer([]string{"192.168.110.131:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer,err:%v\n", err)
		return
	}
	//根据topic取到所有的分区
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { //遍历所有分区
		//针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("ConsumePartition failed,err:%s", err)
			return
		}
		defer pc.AsyncClose()
		//异步从每个分区消费信息
		wg.Add(1)
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d offset:%d key:%s value:%s ", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
		wg.Wait()
	}

}
