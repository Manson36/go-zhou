package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//api参数
	r.GET("/welcome", func(c *gin.Context) {
		//DefaultQuery的第二个参数是默认值
		name := c.DefaultQuery("name", "Jack")
		c.String(http.StatusOK, name)
	})
	r.Run(":8000")
}
