package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/keribend/hello-web/views"
)

var globalCounter int

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", getCounter)
	mux.HandleFunc("POST /", incrementCounter)

	fmt.Println("listening on http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Printf("error listening: %v", err)
	}
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	component := views.Html(globalCounter)
	component.Render(r.Context(), w)
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	globalCounter++
	getCounter(w, r)
}
