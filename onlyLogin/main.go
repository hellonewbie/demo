package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"onlyLogin/api/user"
	"onlyLogin/conf"
	"onlyLogin/utils/handle"
)

func main() {
	//初始化数据
	Load()
	//gin.SetMode(gin.DebugMode) //开发环境
	gin.SetMode(gin.ReleaseMode) //线上环境
	r := gin.Default()
	//r.Use(Auth)
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)
	//r.POST("/logout", user.Logout)
	r.POST("/useVip", user.UseVIPCount)
	r.POST("/Ping", user.Ping)
	users := r.Group("/controlUsers")
	{
		users.GET("/getUsers", user.GetUsers)
		users.POST("/putUsers", user.PutUsers)
		users.POST("/deleteUsers", user.DeleteUsers)
	}
	err := r.Run(":8080")
	if err != nil {
		logrus.Errorf("gin_Engine Run Failed %s", err)
	}
}

func Load() {
	c := conf.Config{}
	c.Routes = []string{"/register", "/login", "/ping", "/useVip"}
	c.OpenJwt = true //开启jwt
	conf.Set(c)
	//初始化数据验证
	handle.InitValidate()
}

//这个用来就是检测是否退出登入，退出后原来的token就不能使用了
//那么登入的时候就会检测一下这个token是否已经在黑名单了
//检测没用cookie就是没有登入
//func Auth(c *gin.Context) {
//	u, err := url.Parse(c.Request.RequestURI)
//	if err != nil {
//		panic(err)
//	}
//	if common.InArrayString(u.Path, &conf.Cfg.Routes) {
//		//next()的方法内部会调用该路由前的其它中间件取执行，然后最后在执行该路由
//		//说白了就是等别的路由完成之后再执行这个路由
//		c.Next()
//		return
//	}
//	if conf.Cfg.OpenJwt {
//		accessToken, has := request.GetParam(c, app.ACCESS_TOKEN)
//		if !has {
//			c.Abort() //判断没有这个参数跳回上一个注册的中间件
//			response.ShowError(c, "nologin")
//			return
//		}
//		ret, err := app.ParseToken(accessToken)
//		if err != nil {
//			c.Abort()
//			response.ShowValidatorError(c, err.Error())
//			return
//		}
//		//函数用于将整型数据转换成指定进制并以字符串的形式返回
//		uid := strconv.FormatInt(ret.UserId, 10)
//		has = app.CheckBlack(uid, accessToken)
//		if has {
//			c.Abort()
//			response.ShowError(c, "nologin")
//			return
//		}
//		c.Set("uid", ret.UserId)
//		c.Next()
//		return
//	}
//	//cookie
//	_, err = c.Cookie(app.COOKIE_TOKEN)
//	if err != nil {
//		c.Abort()
//		response.ShowError(c, "nologin")
//		return
//	}
//	c.Next()
//	return
//}
