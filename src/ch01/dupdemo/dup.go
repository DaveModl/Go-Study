package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//map内置k-v结构，key一般时string,value任意
	//make函数进行map创建
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
		//line := input.Text()
		//counts[line] = counts[line] + 1
	}

	for line, n := range counts{
		if n > 1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}
