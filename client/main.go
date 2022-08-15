package main

import (
	"fmt"
)

func main() {
	var key int
	var loop = true
	for loop {
		fmt.Println("---------------欢迎登录聊天室----------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t   请选择(1-3)")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误,请重新输入")
		}
	}
	//根据用户输入,显示新的信息
	if key == 1 {
		fmt.Println("")

	} else {

	}

}
