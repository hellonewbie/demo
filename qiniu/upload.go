package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	teskSDK "qiniu/GO-SDK/testSDK"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./GO-SDK/template/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		f, err := c.FormFile("f") //从请求中获取携带的参数一样
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})

		} else {
			//将读取的文件保存到本地（服务器本地）
			dst := fmt.Sprintf("./picture/%s", f.Filename)
			c.SaveUploadedFile(f, dst)
			code, url := teskSDK.CoverPicture(f)
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "ok",
				"url":  url,
			})
		}
	})
	r.Run(":8080")
}
