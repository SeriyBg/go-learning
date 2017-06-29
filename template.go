package cms

import (
	"html/template"
)

var Tmpl = template.Must(template.ParseGlob("templates/*"))

type Page struct {
	Title string
	Content string
}
