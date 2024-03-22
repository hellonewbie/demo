package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"onlyLogin/utils/common"
)

//查询数据库所有内容
func GetAllUser() ([]UserPeople, error) {
	users := make([]UserPeople, 0)
	user := new(UserPeople)
	rows, err := mEngine.Where("username >?", 1).Rows(user)
	if err != nil {
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(user)
		if err != nil {
			logrus.Errorf("Rows people failed %s", err)
		}
		users = append(users, *user)
	}
	return users, nil
}

func PutUsers(people UserPeople) (err error) {
	user := new(UserPeople)
	user.Password = common.EBase64(people.Password)
	user.Username = people.Username
	user.VipCount = people.VipCount
	user.StartTime = people.StartTime
	user.EndTime = people.EndTime
	user.Salt = people.Salt
	affected, err := mEngine.Where("username=?", people.Username).Cols("username", "password", "vip_Count", "Start_Time", "end_Time", "salt").Update(user)
	if err != nil {
		logrus.Errorf("Update people failed %s", err)
	}
	fmt.Println(affected)
	return
}

func DeleteUser(people DeletePeople) (Ret bool, err error) {
	user := new(UserPeople)
	user.Username = people.Username
	user.Password = common.EBase64(people.Password)
	affected, err := mEngine.Where("username=? AND password=?", people.Username, user.Password).Delete(user)
	if err != nil {
		logrus.Errorf("DeleteUser failed:%s", err)
		return
	}
	if affected == 0 {
		Ret = false
		return
	} else if affected == 1 {
		Ret = true
		return
	}
	return

}
