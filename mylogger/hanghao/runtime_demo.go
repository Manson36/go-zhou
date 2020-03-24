package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	pc, file, line, ok := runtime.Caller(0) //参数表示向上查找几层函数
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name() //获取函数名
	fmt.Println(funcName)
	fmt.Println(file)            //文件路径+文件名
	fmt.Println(path.Base(file)) //文件名
	fmt.Println(line)            //行号
}
