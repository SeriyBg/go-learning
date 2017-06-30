package main

import (
	"cms"
	"net/http"
)

func main() {
	http.HandleFunc("/", cms.ServerIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.HandleFunc("/page/", cms.ServePage)
	http.HandleFunc("/post/", cms.ServePost)
	http.ListenAndServe(":3000", nil)
}
