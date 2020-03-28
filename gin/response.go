package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

//多种响应方式
func main() {
	r := gin.Default()
	//1.json
	r.GET("/someJson", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJson", "status": 200})
	})
	//2.结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 231

		c.JSON(200, msg)
	})
	//3.XML响应
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})
	//4.YAML响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zhangsan"})
	})
	//5.protobuf格式，谷歌开发的高效存储的工具
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		//定义数据
		label := "label"
		//传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
}
