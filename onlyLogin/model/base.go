package model

import (
	"log"
	"onlyLogin/conf"
	"xorm.io/xorm"
)

var mEngine *xorm.Engine

func init() {
	if mEngine == nil {
		var err error
		mEngine, err = xorm.NewEngine(conf.Db["db1"].DriverName, conf.Db["db1"].Dsn)
		if err != nil {
			log.Fatal(err)
		}
		mEngine.SetMaxIdleConns(conf.Db["db1"].MaxIdle) //空闲连接
		mEngine.SetMaxOpenConns(conf.Db["db1"].MaxOpen) //最大连接数
		mEngine.ShowSQL(conf.Db["db1"].ShowSql)
		mEngine.ShowSQL(true)
	}
}
