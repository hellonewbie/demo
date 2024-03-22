package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.110.131:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v", err)
	}
	defer cli.Close()
	watch := cli.Watch(context.Background(), "name")
	wg.Add(1)
	go func() {
		for wrsp := range watch {
			for _, evt := range wrsp.Events {
				fmt.Printf("type:%s key:%s value:%s", evt.Type, evt.Kv.Key, evt.Kv.Value)
			}
		}
	}()
	wg.Wait()
}
