package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个k-v map计数器
	counts := make(map[string]int)
	//获取用户输入的参数
	files := os.Args[1:]
	if len(files) == 0 {
		//如果不是文件按照标准输入获取数据
		countsLines(os.Stdin,counts)
	}else {
		//range同时返回K-v
		for _, arg := range files{
			//打开文件
			f, err := os.Open(arg)
			if err != nil {
				//如果打开文件出错则输出错误信息
				fmt.Fprintf(os.Stderr,"dup2: %v\n",err)
				continue
			}
			countsLines(f,counts)
			f.Close()
		}
	}
	for line,n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
		
	}

	
}

func countsLines(file *os.File, ints map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		ints[input.Text()]++
	}
	
}
