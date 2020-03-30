package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		//获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			//**给客户端设置cookie**有七个参数
			//1.name:cookie名；2.value；3.maxAge int 单位为秒；4.path cookie所在目录；5.domain string 域名；
			// 6.secure 是否只能通过https访问；7.httpOnly bool 是否有序别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Println("cookie 的值是：", cookie)
	})

	r.Run(":8000")
}
