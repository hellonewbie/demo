package etcd

import (
	"08logagent/common"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	client *clientv3.Client
)

func Init(address []string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed %s", err)
		return
	}
	return
}

//拉取日志收集配置的函数
func GetConf(key string) (collectEntryList []common.CollectEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		logrus.Errorf("get conf from etcd by key :%s failed,err%v", key, err)
		return
	}
	if len(resp.Kvs) == 0 {
		logrus.Warningf("get conf len:0 from etcd by key:%s failed ,err %v", key, err)
		return
	}
	ret := resp.Kvs[0]
	//反序列化后存到collectEntryList中
	err = json.Unmarshal(ret.Value, &collectEntryList)
	if err != nil {
		logrus.Errorf("json unmarshal failed %s", err)
		return
	}
	return
}

func WatchConf(key string) {
	watchCh := client.Watch(context.Background(), key)
	for wresp := range watchCh {
		logrus.Info("get new conf from etcd!")
		for _, evt := range wresp.Events {
			fmt.Printf("type%s key %s value %s \n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}
}
