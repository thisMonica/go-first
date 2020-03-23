//go语言以包作为管理单位
//每个文件必须先声明包
//程序必须有一个main包
package main

import "fmt"

func main() { //左括号必须和函数名同行

	//调用函数，大部分需要先导包
	fmt.Println("Hello World!")
	//go语言语句结束没有分号

	fmt.Printf("1321")
}
