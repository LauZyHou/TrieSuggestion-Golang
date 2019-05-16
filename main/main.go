package main

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func handleSearch(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//设置路由
	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("search", handleSearch)
	//监听端口,如果端口不可用会返回错误,这里对错误进行处理
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("监听失败,出错:%v\n", err)
		return
	}
}
