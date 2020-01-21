package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	a := get(os.Args[1])
	fmt.Println(a)
}

func get(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	return resp
}
