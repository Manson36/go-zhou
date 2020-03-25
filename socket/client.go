package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//与server建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial :20000 failed error:", err)
		return
	}
	//发送数据
	reader := bufio.NewReader(os.Stdin) //os.Stdin 表示输入
	for {
		fmt.Println("请说话")
		msg, _ := reader.ReadString('\n') //读到换行
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}

	conn.Close()
}
