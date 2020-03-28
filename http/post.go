package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://127.0.0.1:9090/post"
	contentType := "application/json"
	data := `{"name":"小王子", "age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed, err:", err)
		return
	}
	fmt.Println(string(b))
}

func post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//1.请求类型是application/x-www-form-urlencoded 时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	//2.请求类型是application/json时从r.Body中获取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request.Body failed, err", err)
		return
	}

	fmt.Println(string(b))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}
