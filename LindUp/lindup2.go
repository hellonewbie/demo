package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//从前端接收的cooie值
type Cookies struct {
	Data string `json:"data"`
}

//包外通过调用这个首字母大写的结构体，间接的调用这个首字母小写的结构体的内容（前提必须是参数首字母大写）
//里面的参数名要和字段名对应起来
type Times struct {
	Time string `form:"time"`
}

type Cookie struct {
	Cookie string `form:"cookie"`
}

type Usetime struct {
	Usetime string `json:"usetime"`
}

//cookie时间长短比较,返回值只需要给定类型即可
func TimeDiff(time1 string, time2 string) int64 {
	loc, _ := time.LoadLocation("Local") //获取时区
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板，必须是这个
	//将日期转为Time格式
	tmp1, _ := time.ParseInLocation(timeLayout, time1, loc)
	tmp2, _ := time.ParseInLocation(timeLayout, time2, loc)
	timestamp1 := tmp1.Unix()         //转化为时间戳 类型是int64
	timestamp2 := tmp2.Unix()         //转化为时间戳 类型是int64
	time := (timestamp2 - timestamp1) //相差的秒数
	fmt.Println(time)                 //输出秒数
	return time
}

//主函数
func main() {
	r := gin.Default()
	Db, err := gorm.Open("mysql", "info:ljy123...@/info?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	Db.SingularTable(true)
	defer Db.Close()
	//有效Cookie存入
	r.POST("/inputcookie", func(ctx *gin.Context) {
		var cookie Cookies
		//接收前端传来的数据
		ctx.BindJSON(&cookie)
		//查询数据条数
		sqlStr := "SELECT COUNT(*) FROM cookie"
		result, err := Db.Debug().DB().Query(sqlStr)
		total := 0
		for result.Next() {
			err := result.Scan(
				&total,
			)
			if err != nil {
				fmt.Println("total make mistake")
			}
		}
		fmt.Println(total)
		if err != nil {
			fmt.Printf("SELECT COOKIE COUNT failed, err: %v\n", err)
			return
		}
		//判断数据条数
		if total != 0 {
			//cookie值比对
			sqlStr := "SELECT cookie FROM cookie"
			rows, err := Db.DB().Query(sqlStr)
			if err != nil {
				fmt.Println("SELECT COOKIE FROM LINDUP ERR")
			}
			var cookies []*Cookie
			//讲查询的内容保存到结构体切片
			for rows.Next() {
				cookie := &Cookie{}
				rows.Scan(&cookie.Cookie)
				cookies = append(cookies, cookie)
			}
			//遍历切片比对cookie值是否已经存在数据库里面
			for i := 0; i < total; i++ {
				fmt.Println(cookies[i].Cookie)
				if cookies[i].Cookie == cookie.Data {
					ctx.JSON(http.StatusOK, gin.H{
						"result": "2",
						"data":   "数据库已有相同数据",
					})
				} else {
					var cookie1 Cookie
					Db.Debug().Model(&cookie1).Update("cookie", cookie.Data)
					//删除时间表的全部数据
					sqlStr := "DELETE FROM times"
					_, err := Db.DB().Exec(sqlStr)
					if err != nil {
						fmt.Println("DELETE FAILED")
					}
					ctx.JSON(http.StatusOK, gin.H{
						"result": "1",
						"data":   "写入成功",
					})
				}
			}

		} else {
			var cookie1 Cookie
			Db.Debug().Model(&cookie1).Update("cookie", cookie.Data)
			//删除时间表的全部数据
			sqlStr := "DELETE FROM times"
			_, err := Db.DB().Exec(sqlStr)
			if err != nil {
				fmt.Println("DELETE FAILED")
			}
			ctx.JSON(http.StatusOK, gin.H{
				"result": "1",
				"data":   "写入成功",
			})
		}

	})
	//比对时间回传cookie
	r.POST("/outputcookie", func(ctx *gin.Context) {
		//到期时间
		var useTime Usetime
		ctx.BindJSON(&useTime)
		sqlStr := "SELECT COUNT(*) FROM times"
		result, err := Db.Debug().DB().Query(sqlStr)
		//数据条数
		total := 0
		for result.Next() {
			err := result.Scan(
				&total,
			)
			if err != nil {
				fmt.Println("total make mistake")
			}
		}
		fmt.Println(total)
		if err != nil {
			fmt.Printf("SELECT TIMES COUNT failed, err: %v\n", err)
			return
		}
		//调用时间比对函数
		//用来遍历的结构体切片
		var time1s []*Times
		//保存当前时间的结构体变量
		var time2 Times
		sqlStr1 := "SELECT created FROM times"
		rows, _ := Db.DB().Query(sqlStr1)
		for rows.Next() {
			time1 := &Times{}
			err := rows.Scan(&time1.Time)
			time1s = append(time1s, time1)
			if err != nil {
				fmt.Println("SELECT CREATE FORM TIMES FAILED")
			}
		}
		//用于数据库查询
		//查出发出请求的时间
		sqlStr2 := "SELECT NOW()"
		row2 := Db.DB().QueryRow(sqlStr2)
		row2.Scan(&time2.Time)
		time4, _ := time.Parse(time.RFC3339, time2.Time)
		//用来遍历保存时间
		var usedtime [2]int64
		for i := 0; i < len(time1s); i++ {
			fmt.Println(time1s[i])
			time3, _ := time.Parse(time.RFC3339, time1s[i].Time)
			usedtime[i] = TimeDiff(time4.Format("2006-01-02 15:04:05"), time3.Format("2006-01-02 15:04:05"))
			println("第", i, usedtime[i])
		}
		if total < 2 {
			sqlStr = "INSERT INTO times(created) VALUES (?)"
			_, err := Db.DB().Exec(sqlStr, useTime.Usetime)
			if err != nil {
				fmt.Println("INSERT INTO FAILED")
			}
			sqlStr := "select * from  cookie "
			rows, err := Db.DB().Query(sqlStr)
			fmt.Println(rows)
			if err != nil {
				fmt.Println(err)
			}
			var infos []*Cookie
			for rows.Next() {
				info := &Cookie{}
				rows.Scan(&info.Cookie)
				infos = append(infos, info)
			}
			ctx.JSON(http.StatusOK, gin.H{
				"result": "1",
				"info":   infos,
			})
		}
		for i := 0; i < 2; i++ {
			if total >= 2 && usedtime[i] < 0 {
				var timemin Times //为了用来接收最小时间值
				sqlStr1 := "SELECT created FROM times WHERE created= (SELECT MIN(created) FROM times)"
				row := Db.DB().QueryRow(sqlStr1)
				row.Scan(&timemin.Time)
				time3, _ := time.Parse(time.RFC3339, timemin.Time)
				//用于数据库查询
				time5 := time3.Format("2006-01-02 15:04:05")
				fmt.Println(time5)
				sqlStr3 := fmt.Sprintf("UPDATE times SET created='%s' WHERE created='%s'", useTime.Usetime, time5)
				fmt.Println(sqlStr3)
				_, err := Db.DB().Exec(sqlStr3)
				if err != nil {
					fmt.Println("UPDATA TIMES FAILED")
				}
				sqlStr := "select * from  cookie "
				rows, err := Db.DB().Query(sqlStr)
				fmt.Println(rows)
				if err != nil {
					fmt.Println(err)
				}
				var infos []*Cookie
				for rows.Next() {
					info := &Cookie{}
					rows.Scan(&info.Cookie)
					infos = append(infos, info)
				}
				ctx.JSON(http.StatusOK, gin.H{
					"result": "11",
					"data":   infos,
				})
				break
			}
		}
		if total >= 2 && usedtime[0] > 0 && usedtime[1] > 0 {
			var timemin Times //为了用来接收最小时间值
			sqlStr1 := "SELECT created FROM times WHERE created= (SELECT MIN(created) FROM times)"
			row := Db.DB().QueryRow(sqlStr1)
			row.Scan(&timemin.Time)
			time3, _ := time.Parse(time.RFC3339, timemin.Time)
			//用于数据库查询
			time5 := time3.Format("2006-01-02 15:04:05")
			ctx.JSON(http.StatusOK, gin.H{
				"result": "2",
				"data":   time5,
			})
		}

	})
	r.POST("/dropall", func(ctx *gin.Context) {
		sqlStr6 := "DELETE FROM times"
		_, err := Db.DB().Exec(sqlStr6)
		if err != nil {
			fmt.Println("DELETE TABLE FAILED")
		}
	})

	r.Run(":12311")
}
