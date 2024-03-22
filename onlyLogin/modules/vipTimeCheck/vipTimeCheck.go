package vipTimeCheck

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

//cookie时间长短比较,返回值只需要给定类型即可
func TimeDiff(time1 string, temp2 time.Time) int64 {
	loc, err := time.LoadLocation("Local") //获取时区
	if err != nil {
		logrus.Errorf("Time_Location get failed %s", err)
	}
	timeLayout := "2006-01-02 15:04:05" //转化所需模板，必须是这个
	//将日期转为Time格式
	tmp1, err := time.ParseInLocation(timeLayout, time1, loc)
	timestamp1 := tmp1.Unix()         //转化为时间戳 类型是int64
	timestamp2 := temp2.Unix()        //转化为时间戳 类型是int64
	time := (timestamp2 - timestamp1) //相差的秒数
	fmt.Println(time)                 //输出秒数
	return time
}
