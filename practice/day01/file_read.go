package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile() {
	fileObj, err := os.Open("./practice/day01/file_read.go")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer fileObj.Close()

	var tmp = make([]byte, 128)
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}

		fmt.Println("读取了", n)
		fmt.Println(string(tmp[:]))
	}
}

func readFileByBufio() {
	fileObj, err := os.Open("./practice/day01/file_read.go")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			return
		}

		if err != nil {
			fmt.Println("read file error", err)
			return
		}

		fmt.Println(line)
	}
}

func readFileByioutil() {
	ret, err := ioutil.ReadFile("./practice/day01/file_read.go")
	if err != nil {
		fmt.Println("读取文件失败， error:", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	//readFile()
	//readFileByBufio()
	readFileByioutil()
}
