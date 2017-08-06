package main

import (
	"fmt"
	"os"
) 

func main() {
	for index,arg := range os.Args {
		fmt.Print("index: ")
		fmt.Print(index)
		fmt.Print(",arg: ")
		fmt.Println(arg)
	}   
}
