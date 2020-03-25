package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//str := `<h1 style="color:red"> hello, 女神！</h1>`
	//从文件中读取
	str, err := ioutil.ReadFile("./http/xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("没有内容")))
	}
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) {
	//对于GET请求，参数都在url上，请求体是没有数据的。
	fmt.Println(r.URL)
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello", f1)
	http.HandleFunc("/xxx", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
