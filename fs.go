package main

import (
	"os"
	"net/http"
	"flag"
)

func main() {
	var dir string
	port := flag.String("port", "3000", "Port to start HTTP on.")
	path := flag.String("path", "", "Path to server")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}

	http.ListenAndServe(":" + *port, http.FileServer(http.Dir(dir)))
}
