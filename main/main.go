package main

import (
	"fmt"
	"github.com/unrolled/render"
	"net/http"
)

//全局变量
var (
	renderHTML *render.Render
)

//初始化的函数,在这里创建渲染HTML用的render对象
func init() {
	//给render用的的选项对象
	option := render.Options{
		//指定模板文件的目录位置
		Directory: "views",
		//指定模板文件的扩展名
		Extensions: []string{".tmpl", ".html"},
	}
	//render渲染对象
	renderHTML = render.New(option)
}

//访问/index时,渲染返回一个HTML页面
func handleIndex(w http.ResponseWriter, r *http.Request) {
	//用render对象渲染HTML页面,参数(响应对象,状态码,模板文件名,模板绑定的变量)
	renderHTML.HTML(w, http.StatusAccepted, "index", nil)
}

//访问/search时,返回JSON数据
func handleSearch(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//初始化(将数据之类的获取到,预处理好等过程)
	err := Init()
	if err != nil {
		fmt.Printf("初始化失败:%v\n", err)
		return
	}

	//设置路由
	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("/search", handleSearch)
	//监听端口,如果端口不可用会返回错误,这里对错误进行处理
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("监听失败,出错:%v\n", err)
		return
	}
}
