package main

import (
	"../logic"
	"fmt"
)

//服务器初始化过程,所有小的初始化方法都在这里调用
func Init() (err error) {
	err = logic.InitLogic("./data/school.dat")
	if err != nil {
		return
	}
	fmt.Print("初始化成功\n")
	return
}
