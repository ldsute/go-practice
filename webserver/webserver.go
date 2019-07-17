package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Welcome to my web server\n")
	})
	http.ListenAndServe(":8080", nil)
}
