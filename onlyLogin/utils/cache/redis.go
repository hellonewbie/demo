package cache

import (
	"github.com/gomodule/redigo/redis"
	"onlyLogin/conf"
	"time"
)

var RedisClient *redis.Pool
var RedisExpire = 86400 * 7
var RedisSuf = "onlyLogin.cache.redis."

func init() {
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,                 //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 * time.Second, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			c, err := redis.Dial(conf.Redis["type"], conf.Redis["address"])
			if err != nil {
				return nil, err
			}
			if conf.Redis["auth"] != "" {
				if _, err := c.Do("AUTH", conf.Redis["auth"]); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
}
