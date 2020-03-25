package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//resp, err := http.Post("http://127.0.0.1:9090/xxx", "image/jpeg", &buf)
	resp, err := http.Get("http://127.0.0.1:9090/xxx?name=liqin&age=29")
	if err != nil {
		fmt.Println("get url failed, err:", err)
		return
	}

	defer resp.Body.Close() //response使用之后一定记得关闭
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body err:", err)
		return
	}
	fmt.Println(string(b))
}
