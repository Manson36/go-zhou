package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	//3.监听端口，默认是8080
	r.Run(":8000")
}
