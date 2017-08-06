package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	x, _ := strconv.Atoi(input.Text())
	fmt.Println(Signum(x))
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default: 
		return 0
	case x < 0:
		return -1
	}
}
