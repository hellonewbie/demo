package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type UserRow struct {
	Username int64  `json:"username" `
	VipCount int64  `json:"vipCount" `
	EndTime  string `json:"endTime" `
}
type UserPeople struct {
	Username  int64  `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	VipCount  int64  `json:"vipCount"`
	StartTime string `json:"startTime" binding:"required"`
	EndTime   string `json:"endTime"  binding:"required"`
	Salt      string `json:"salt" binding:"required"`
}

type DeletePeople struct {
	Username int64  `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Salt     string `json:"salt" xorm:"not null comment('盐值') CHAR(4)"`
}

//var UsersStatusOk = 1
//var UsersStatusDel = 10
//var UsersStatusDef = 0
//
//var usersTable = "usercount"

func (u *UserPeople) Register() bool {
	affected, err := mEngine.Insert(u)
	if err == nil {
		fmt.Println(affected)
		return true
	}
	return false
}

func (u *UserPeople) GetRow() bool {
	//从数据库中查一条数据
	has, err := mEngine.Get(u)
	if err == nil && has {
		return true
	}
	return false
}

func (u *UserPeople) GetUser(Id int64) (UserPeople, bool) {
	user := new(UserPeople)
	rows, err := mEngine.Where("username=?", Id).Rows(user)
	if err != nil {
		return *user, false
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(user)
	}
	return *user, true
}

//更新超级VIP使用次数
func (u *UserPeople) UsVipCount(Id int64) bool {
	user := new(UserPeople)
	users, Bool := user.GetUser(Id)
	fmt.Println(users.VipCount)
	if Bool {
		if users.VipCount > 0 {
			users.VipCount = users.VipCount - 1
			user.VipCount = users.VipCount
			fmt.Println(user.VipCount)
		} else {
			return false
		}
	}
	//这里有个坑，xorm不会提取结构体里面的零值，所以把要更新的列名加入之后就不会出现这样的问题
	_, err := mEngine.Where("username=?", Id).Cols("vip_Count").Update(user)
	if err != nil {
		logrus.Errorf("UsVipCount update failed %s", err)
	}
	return true
}
