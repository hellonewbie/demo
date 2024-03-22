package main

import (
	"08logagent/etcd"
	"08logagent/kafka"
	"08logagent/tailfile"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

//日志收集的客户端
//类似的开源项目还有filebeat
//收集指定目录下的日志文件，发送到kafka中

//整个LogAgent的配置结构体
type Config struct {
	KafkaConfig `ini:"kafka"`
	Collect     `ini:"collect"`
	EtcdConfig  `ini:"etcd"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	Topic    string `ini:"topic"`
	ChanSize int64  `ini:"chan_size"`
}

type Collect struct {
	LogFilePath string `ini:"logfile_path"`
}

type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"`
}

func main() {
	var configObj = new(Config)
	//0.初始化（读配置文件`ini`）
	//cfg, err := ini.Load("./conf/config.ini")
	//if err != nil {
	//	logrus.Error("load config failed ,err :%v", err)
	//	return
	//}
	//kafka := cfg.Section("kafka").Key("address").String()
	//fmt.Println(kafka)
	//MapTo将数据源映射到给定的结构。
	err := ini.MapTo(configObj, "./conf/config.ini")
	if err != nil {
		logrus.Error("load config failed,err %v", err)
		return
	}
	fmt.Printf("%#v", configObj)
	logrus.Info("ini MapTo success")
	fmt.Println(configObj.KafkaConfig.ChanSize)
	err = etcd.Init([]string{configObj.EtcdConfig.Address})
	if err != nil {
		logrus.Errorf("etcd init failed %s", err)
		return
	}
	logrus.Info("init ectd success")

	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Error("kafka Init failed %v", err)
		return
	}
	logrus.Info("init kafka success")

	allConf, err := etcd.GetConf(configObj.EtcdConfig.CollectKey)
	if err != nil {
		logrus.Errorf("Get conf from etcd failed%s", err)
	}
	fmt.Println(allConf)
	go etcd.WatchConf(configObj.EtcdConfig.CollectKey)
	//1.读配置文件
	//2.根据配置中的路径使用tail去收集日志
	//3.把日志通过sarama发往kafka
	err = tailfile.Init(allConf)
	if err != nil {
		logrus.Error("tailfile init failed :%v", err)
		return
	}
	logrus.Info("init tailFile success")
}
