package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"onlyLogin/model"
	"onlyLogin/modules/app"
	"onlyLogin/modules/vipTimeCheck"
	"onlyLogin/utils/common"
	"onlyLogin/utils/response"
	"strconv"
	"time"
)

type UsersUsernamePassword struct {
	Username int64  `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var UserRegister model.UserPeople
	if err := c.BindJSON(&UserRegister); err != nil {
		msg := "注册参数输入不完整请检查：" +
			"1.username(int64)" +
			"2.password(string)" +
			"3.vipCount(int64)" +
			"4.startTime(string)" +
			"5.endTime(string)" +
			"6.salt(string)"
		response.ShowValidatorError(c, msg)
		return
	}
	models := model.UserPeople{
		Username:  UserRegister.Username,
		Password:  common.EBase64(UserRegister.Password),
		VipCount:  UserRegister.VipCount,
		StartTime: UserRegister.StartTime,
		EndTime:   UserRegister.EndTime,
		Salt:      UserRegister.Salt,
	}
	if has := models.Register(); !has {
		response.ShowError(c, "registerFailed")
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "注册成功",
		"count": UserRegister,
	})

}

func Login(c *gin.Context) {
	var UserPeople UsersUsernamePassword
	//判断参数是否绑定成功
	if err := c.BindJSON(&UserPeople); err != nil {
		msg := "登入失败，账号或密码错误"
		response.ShowValidatorError(c, msg)
		return
	}
	//判断用户是否存在
	models := model.UserPeople{Username: UserPeople.Username}
	if has := models.GetRow(); !has {
		response.ShowError(c, "user_error")
		return
	}
	//调试代码
	fmt.Println(common.EBase64(UserPeople.Password))
	//判断密码是否正确
	if common.EBase64(UserPeople.Password) != models.Password {
		response.ShowError(c, "login_error")
		return
	}

	//绑定成功之后设置cooke、和token
	ReturnToken, err := app.DoLogin(c, models)
	if err != nil {
		response.ShowError(c, "fail")
	}

	//查询数据库用户信息返回
	var User model.UserPeople
	var has bool
	User, has = models.GetUser(models.Username)
	//回来再改 直接用原生

	timeStr := time.Now()
	RmTime := vipTimeCheck.TimeDiff(User.EndTime, timeStr)
	if err != nil {
		logrus.Errorf("TimeDiff failed %s", err)
	}
	var msg string
	if RmTime >= 0 && has {
		msg = "账号会员已过期请联系管理员进行续费"
	}
	if has {
		c.JSON(200, gin.H{
			"code":         200,
			"msg":          "登录成功",
			"username":     User.Username,
			"endTime":      User.EndTime,
			"vipCount":     User.VipCount,
			"Note":         msg,
			"salt":         User.Salt,
			"access-token": ReturnToken,
		})
	} else {
		msg := "请重新登入"
		response.ShowValidatorError(c, msg)
		return
	}
	return
}

//退出登入
//func Logout(c *gin.Context) {
//	//access_token  refresh_token 加黑名单
//	accessToken, has := request.GetParam(c, app.ACCESS_TOKEN)
//	if has {
//		uid, err := c.Cookie(app.COOKIE_TOKEN)
//		if err != nil {
//			logrus.Errorf("Logout Get Cookie failed %s", err)
//		}
//		//uid := strconv.FormatInt(c.MustGet("UserId").(int64), 10)
//		err = app.AddBlack(uid, accessToken)
//		if err != nil {
//			logrus.Errorf("AddBlack failed :%s", err)
//		}
//		err = app.DelWhite(uid)
//		if err != nil {
//			logrus.Errorf("DelWhite token failed :%s", err)
//		}
//	}
//
//	c.SetCookie(app.COOKIE_TOKEN, "", -1, "/", "", false, true)
//	c.SetCookie(app.ACCESS_TOKEN, "", -1, "/", "", false, true)
//	c.SetCookie(app.REFRESH_TOKEN, "", -1, "/", "", false, true)
//	response.ShowSuccess(c, "success")
//	return
//}

type CheckToken struct {
	AccessToken string `json:"access-token" binding:"required"`
}

//可以用来检测是否登入成功
func Ping(c *gin.Context) {
	var checkwhite CheckToken
	if err := c.BindJSON(&checkwhite); err != nil {
		msg := "Token获取失败，请注意参数名为access-token"
		response.ShowValidatorError(c, msg)
		return
	}
	//accessToken, has := request.GetParam(c, app.ACCESS_TOKEN)
	ret, err := app.ParseToken(checkwhite.AccessToken)
	if err != nil {
		c.Abort()
		response.ShowValidatorError(c, err.Error())
		return
	}
	uid := strconv.FormatInt(ret.UserId, 10)
	//查询Vip次数
	var UserInfo model.UserPeople
	UserInfo.Username = ret.UserId
	UserInfo.GetRow()
	if app.CheckWhite(uid, checkwhite.AccessToken) {
		c.JSON(200, gin.H{
			"code":     200,
			"msg":      "登入成功",
			"end-time": UserInfo.EndTime,
			"vipCount": UserInfo.VipCount,
			"salt":     UserInfo.Salt,
		})
	} else {
		response.ShowError(c, "reLogin")
		return
	}

}

//使用超级VIP权限
func UseVIPCount(c *gin.Context) {
	var checkwhite CheckToken
	var checkModel model.UserPeople
	if err := c.BindJSON(&checkwhite); err != nil {
		msg := "Token获取失败，请注意参数名为access-token"
		response.ShowValidatorError(c, msg)
		return
	}
	ret, err := app.ParseToken(checkwhite.AccessToken)
	if err != nil {
		c.Abort()
		response.ShowValidatorError(c, err.Error())
		return
	}
	//int64转 string类型
	uid := strconv.FormatInt(ret.UserId, 10)
	if app.CheckWhite(uid, checkwhite.AccessToken) {
		//然后就执行更新
		if checkModel.UsVipCount(ret.UserId) {
			var User model.UserPeople
			User, _ = User.GetUser(ret.UserId)
			c.JSON(200, gin.H{
				"code":     200,
				"msg":      "超级权限使用成功",
				"username": User.Username,
				"endTime":  User.EndTime,
				"vipCount": User.VipCount,
			})
		} else {
			response.ShowValidatorError(c, "超级权限已全部用完，需要请联系管理员")
			return
		}
	} else {
		response.ShowValidatorError(c, "Token parse failed,请重新登入")
		return
	}

}

//获取所有用户
func GetUsers(c *gin.Context) {
	users, err := model.GetAllUser()
	if err != nil {
		response.ShowValidatorError(c, "获取失败,请联系后端管理员")
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"users": users,
	})
}

//更新用户信息
func PutUsers(c *gin.Context) {
	var PutUsers model.UserPeople
	if err := c.BindJSON(&PutUsers); err != nil {
		msg := "更新参数输入不完整请检查：" +
			"1.username(int64)" +
			"2.password(string)" +
			"3.vipCount(int64)" +
			"4.startTime(string)" +
			"5.endTime(string)" +
			"6.salt(string)"
		response.ShowValidatorError(c, msg)
		return
	}
	err := model.PutUsers(PutUsers)
	if err != nil {
		logrus.Errorf("PutUsers err:%s", err)
		return
	}
	c.JSON(200, gin.H{
		"code":     200,
		"ReSetMsg": PutUsers,
	})
}

//删除用户信息
func DeleteUsers(c *gin.Context) {
	var DeleteUser model.DeletePeople
	if err := c.BindJSON(&DeleteUser); err != nil {
		msg := "删除前需要比对用户名和密码，请正确输入1.username 2.password 3.salt"
		response.ShowValidatorError(c, msg)
		return
	}
	Ret, err := model.DeleteUser(DeleteUser)
	if err != nil {
		logrus.Errorf("DeleteUsers failed :%s", err)
		return
	}
	if !Ret {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户名或密码输出错误",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "用户删除成功",
		})
	}

}
