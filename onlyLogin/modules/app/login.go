package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"onlyLogin/conf"
	"onlyLogin/model"
	"strconv"
	"time"
)

func DoLogin(c *gin.Context, user model.UserPeople) (string, error) {
	var ReturnToken string
	if conf.Cfg.OpenJwt {
		customClaims := &CustomClaims{
			UserId:  user.Username,
			EndTime: user.EndTime,
			//获取请求ip
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Duration(MAXAGE) * time.Second).Unix(), //过期时间
			},
		}
		accessToken, err := customClaims.MakeToken()
		ReturnToken = accessToken
		if err != nil {
			return "", err
		}
		//refreshClaims := &CustomClaims{
		//	UserId:  user.Username,
		//	EndTime: user.EndTime,
		//	Ip:      c.ClientIP(),
		//	StandardClaims: jwt.StandardClaims{
		//		ExpiresAt: time.Now().Add(time.Duration(MAXAGE+1800) * time.Second).Unix(), //过期时间
		//	},
		//}
		//refreshToken, err := refreshClaims.MakeToken()
		//if err != nil {
		//	return "", err
		//}
		//c.Header(ACCESS_TOKEN, accessToken)
		//c.Header(REFRESH_TOKEN, refreshToken)
		//c.SetCookie(ACCESS_TOKEN, accessToken, MAXAGE, "/", "", false, true)
		//c.SetCookie(REFRESH_TOKEN, refreshToken, MAXAGE, "/", "", false, true)
		//将登入后的cookie设置到白名单中
		err = AddWhite(strconv.Itoa(int(user.Username)), accessToken)
		if err != nil {
			logrus.Errorf("AddWhite Failed %s", err)
		}
	}
	//id := strconv.Itoa(int(user.Username))
	//c.SetCookie(COOKIE_TOKEN, id, MAXAGE, "/", "", false, true)
	return ReturnToken, nil
}
