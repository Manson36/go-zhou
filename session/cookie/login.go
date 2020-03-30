package main

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		//返回错误
		c.JSON(401, gin.H{"error": "err"})
		//若验证不通过，不再调用后面的函数处理；如果不加c.Abort()，会输出"data":"home".
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		//设置cookie
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		//返回信息
		c.String(200, "Login success")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	r.Run(":8000")
}
