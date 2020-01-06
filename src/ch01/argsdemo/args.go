package main

import (
	"fmt"
	"os"
)

func main() {
	//声明变量，两个字符串---字符串定义是反过来的
	var s, sep string
	//for循环不需要括号
	for i := 0; i < len(os.Args); i++  {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	//打印k-v
	str := " "
	for kArgs, vArgs := range os.Args[0:] {
		fmt.Print(kArgs)
		fmt.Print(str)
		fmt.Println(vArgs)
	}
}
/**
几种for循环形式
这三部分均可省略
for initialization;condition;post{
}
while
for condition{
}
死循环
for{
}

 */