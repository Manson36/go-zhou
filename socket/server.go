package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	//本地端口启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed error:", err)
		return
	}
	//等待别人跟我建立链接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go process(conn) //启动一个goroutine处理链接
	}
}

func process(conn net.Conn) {
	defer conn.Close() //关闭链接
	//客户端通信
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed error", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) //发送数据
	}
}
