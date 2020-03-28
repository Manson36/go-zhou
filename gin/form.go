package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	//binding:"required" 修饰的字段，若接收值为空，则报错，是必须字段
	User     string `json:"username" form:"username" uri:"user" xml:"user" binding:"required"`
	Password string `json:"password" form:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/loginForm", func(c *gin.Context) {
		//声明接收的变量
		var form Login
		//Bind()默认解析并绑定form格式
		//根据请求头中content-type自动判断
		if err := c.Bind(&form); err != nil {
			//返回错误信息
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//需要html去渲染才能取得返回值，这里先使用json返回
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.Run(":8000")
}
