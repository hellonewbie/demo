package TimeDiff

import (
	"fmt"
	"time"
)

//cookie时间长短比较,返回值只需要给定类型即可
func TimeDiff(time1 string, time2 string) int64 {
	loc, _ := time.LoadLocation("Local") //获取时区
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板，必须是这个
	//将日期转为Time格式
	tmp1, _ := time.ParseInLocation(timeLayout, time1, loc)
	tmp2, _ := time.ParseInLocation(timeLayout, time2, loc)
	timestamp1 := tmp1.Unix() //转化为时间戳 类型是int64
	timestamp2 := tmp2.Unix() //转化为时间戳 类型是int64

	time := (timestamp2 - timestamp1) //相差的秒数
	fmt.Println(time)                 //输出秒数
	return time
}
