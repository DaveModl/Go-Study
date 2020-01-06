package main

import (
	"fmt"
	"os"
)

func main() {
	//初始化两个字符串
	s, sep := "",""
	//空白标识符，接收一些不想处理的值,这里是不处理key
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
/**
go不允许声明变量之后后续不进行使用，这是一个编译错误
声明变量的方式---一般使用前两种
s := ""
var s string
var s = ""
var s string = ""
 */
