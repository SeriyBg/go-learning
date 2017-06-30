package cms

import (
	"net/http"
	"time"
	"strings"
)

func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)
	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()
		if contentType == "page" {
			Tmpl.ExecuteTemplate(w, "page", &Page{
				Title:   title,
				Content: content,
			})
			return
		}
		if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "post", &Post{
				Title:   title,
				Content: content,
			})
			return
		}
	default:
		http.Error(w, "Method not supported: "+r.Method, http.StatusMethodNotAllowed)
	}

}

func ServerIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Go Projects CMS",
		Content: "Welcome to out home page!",
		Posts: []*Post{
			&Post{
				Title:         "Hello, World!",
				Content:       "Hello World! Very good content!",
				DatePublished: time.Now(),
			},
			&Post{
				Title:         "Second post!",
				Content:       "Still very nice content",
				DatePublished: time.Now().Add(-time.Hour),
				Comments: []*Comment{
					&Comment{
						Author:        "Doe John",
						Comment:       "I'm impressed!",
						DatePublished: time.Now().Add(-time.Hour / 2),
					},
				},
			},
		},
	}
	Tmpl.ExecuteTemplate(w, "page", p)
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/page/")
	if path == "" {
		http.NotFound(w, r)
		return
	}
	p := &Page{
		Title: strings.ToTitle(path),
		Content: "Here is my page",
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}

func ServePost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/post/")
	if path == "" {
		http.NotFound(w, r)
		return
	}

	p := &Post{
		Title: strings.ToTitle(path),
		Content: "Here is my post",
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}
