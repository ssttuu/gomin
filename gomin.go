package main

import (
	"net/http"
	"github.com/stupschwartz/gomin/minify"
)

func main() {
	http.HandleFunc("/", minify.RootHandler)
	http.ListenAndServe(":8080", nil)
}
