package main

import (
	"cms"
	"os"
)

func main() {
	p := &cms.Page{
		Title: "Hello, World!",
		Content: "Body of the web page",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
