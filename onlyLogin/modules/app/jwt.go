package app

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"onlyLogin/utils/cache"
)

const (
	SECRETKEY         = "e1c5179f5b2c277a"
	MAXAGE            = 3600 * 24
	CACHE_BLACK_TOKEN = "black.token."
	CACHE_White_TOKEN = "white.token."
)

type CustomClaims struct {
	UserId  int64
	EndTime string
	Ip      string
	jwt.StandardClaims
}

//产生token

func (cc *CustomClaims) MakeToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	return token.SignedString([]byte(SECRETKEY))
}

//解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
	// 要传入指针，项目中结构体都是用指针传递，节省空间。
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

//检查token是否在白名单中
func CheckWhite(key string, token string) bool {
	key = cache.RedisSuf + CACHE_White_TOKEN + key
	//从连接池里获取连接
	rc := cache.RedisClient.Get()
	//用完将连接放回连接池
	defer rc.Close()
	val, err := redis.String(rc.Do("GET", key))
	if err != nil || val != token {
		return false
	}
	return true
}

//检查token是否在黑名单
func CheckBlack(key, token string) bool {
	key = cache.RedisSuf + CACHE_BLACK_TOKEN + key
	//从连接池里获取连接
	rc := cache.RedisClient.Get()
	//用完将连接放回连接池
	defer rc.Close()
	val, err := redis.String(rc.Do("GET", key))
	if err != nil || val != token {
		return false
	}
	return true
}

//加入到白名单中  //这有的黑名单不能有
func AddWhite(key, token string) (err error) {
	key = cache.RedisSuf + CACHE_White_TOKEN + key
	rc := cache.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	_, err = rc.Do("Set", key, token, "EX", MAXAGE)
	if err != nil {
		return
	}
	return
}

//加入到黑名单中
func AddBlack(key, token string) (err error) {
	key = cache.RedisSuf + CACHE_BLACK_TOKEN + key
	// 从池里获取连接
	rc := cache.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	_, err = rc.Do("Set", key, token, "EX", MAXAGE)
	if err != nil {
		return
	}
	return
}
func DelWhite(key string) (err error) {
	key = cache.RedisSuf + CACHE_White_TOKEN + key
	rc := cache.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	_, err = rc.Do("del", key)
	if err != nil {
		return
	}
	return
}
