package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//ioutil读取文件内容
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		//每行内容以换行符分隔
		for _,line := range strings.Split(string(data),"\n"){
			counts[line]++
		}

		for line,n := range counts{
			if n > 1 {
				fmt.Printf("%d\t%s\n",n,line)
			}

		}

	}
}