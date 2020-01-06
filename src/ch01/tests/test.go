package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)!=0{
		fmt.Println(os.Args[0])
	}
	fmt.Println(os.Args[1])
}
