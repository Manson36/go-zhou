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

	r.GET("/:user/:password", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.Run(":8000")
}
