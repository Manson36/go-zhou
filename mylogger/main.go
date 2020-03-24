package main

import (
	"github.com/go-zhou/mylogger/mylog"
	"time"
)

func main() {
	//测试我们写的日志
	//log := mylog.NewLogger
	log := mylog.Logger{}
	for {
		log.Debug("这是一条debug日志")
		log.Info("这是一条ingo日志")
		time.Sleep(time.Second)
	}

}
