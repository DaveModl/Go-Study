package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//连接字符串,相对高效
	fmt.Println(strings.Join(os.Args[1:]," "))
}
