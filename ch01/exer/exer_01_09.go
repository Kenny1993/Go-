package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var prefix string = "http://"

func main() {
	for _, url := range os.Args[1:] {
		fmt.Println("before adding prefix: " + url)
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		fmt.Println("after adding prefix: " + url)			
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Print("http status code: ")
		fmt.Println(resp.Status)
		byteLen, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d", byteLen)
	}
}
