package main

import (
	"fmt"
	"os"
) 

func main() {
	var s,sep string
	len := len(os.Args)
	for i := 0; i < len; i++ {
		s += sep + os.Args[i]
		sep = " "
	}   
	fmt.Println(s)
}
