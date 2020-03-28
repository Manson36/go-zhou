package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20

	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		//获取所有文件
		file := form.File["file"]
		//遍历所有文件
		for _, f := range file {
			//逐个存
			if err := c.SaveUploadedFile(f, f.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}

		//打印信息
		c.String(http.StatusOK, fmt.Sprintf("upload ok %d file", len(file)))
	})
	r.Run(":8000")
}
