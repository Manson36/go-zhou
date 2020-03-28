package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//表单参数
	r.POST("/form", func(c *gin.Context) {
		//表单参数设置默认值
		type1 := c.DefaultPostForm("type", "alert")
		//接收其他的
		username := c.PostForm("username")
		password := c.PostForm("password")
		//多选框
		hobbys := c.PostFormArray("hobby")

		c.String(http.StatusOK,
			fmt.Sprintf("type is %s, username is %s, password is %s, hobbys is %v",
				type1, username, password, hobbys))
	})

	r.Run(":8000")
}
