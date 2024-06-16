package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/keribend/hello-web/views"
)

func main() {
	http.Handle("/", templ.Handler(views.Html()))

	http.ListenAndServe(":8080", nil)
}
