package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		//设置变量到Context的key中，可以通过Get()获取
		c.Set("request", "中间件")
		//执行函数，c.Next()之后才是执行下面代码的内容
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	r := gin.Default()
	//注册中间件
	r.Use(MiddleWare())
	//加{} 是为了代码规范
	{
		r.GET("/middleware", func(c *gin.Context) {
			//取值
			req, _ := c.Get("request")
			fmt.Println("request", req)
			//页面接收
			c.JSON(200, gin.H{"request": req})
		})
		//根路由后面是定义的局部中间件
		r.GET("middleware2", MiddleWare(), func(c *gin.Context) {
			//取值
			req, _ := c.Get("request")
			fmt.Println("request", req)
			c.JSON(200, gin.H{"request": req})
		})
	}

	r.Run(":8000")
}
