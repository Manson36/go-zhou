package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//API参数
	r.GET("/xxx/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, name)
	})
	r.Run(":8000")
}
