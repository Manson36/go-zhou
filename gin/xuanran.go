package main

import "github.com/gin-gonic/gin"

//HTML渲染
func main() {
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("template/*") //template目录下的所有文件
	//r.LoadHTMLFiles("template/index.tmpl")//模板文件名称：template目录下的index.tmpl文件
	r.GET("/index", func(c *gin.Context) {
		//根据文件名渲染
		//最终json将title替换
		c.HTML(200, "index.html", gin.H{"title": "我的标题"})
	})

	r.Run(":8000")
}
