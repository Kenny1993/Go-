package main

import (
	"log"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/download/go.tar.gz", handler)
	log.Fatal(http.ListenAndServe("11.11.1.10:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("/usr/local/go1.8.3.linux-amd64.tar.gz")
	w.Write(data)
}
