package main

import (
	"fmt"
	"os"
	"time"
	"strings"
) 

func main() {
	args := os.Args
	s1,sep := "",""
	before1 := time.Now()
	for _,arg := range args {
		s1 += sep + arg
		sep = " "
	}   
	fmt.Print("without strings.Join: ")
	fmt.Println(time.Now().Sub(before1))
	fmt.Println(s1)
	before2 := time.Now()
	s2 := strings.Join(args," ")	
	fmt.Print("with strings.Join: ")
	fmt.Println(time.Now().Sub(before2))
	fmt.Println(s2)
}
