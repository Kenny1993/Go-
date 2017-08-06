package main

import (
	"fmt"
	"bufio"
	"os"
)

var (
	heads = 0
	tails = 0
)

func main() {
	switch conflip() {
		case "heads":
			heads++
		case "tails":
			tails++
		default:
			fmt.Println("landed on edge!")
	}
	if heads > 0 {
		fmt.Print("heads")
	}
	if tails > 0 {
		fmt.Print("tails")
	}
}

func conflip() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}
