package main

import (
	// "fmt"
	"html/template"
	"net/http"

	"go.uber.org/zap"
)

type WebEngine struct {
	log *zap.Logger
}

func NewWebEngine() *WebEngine {
	e := WebEngine{}
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	e.log = log
	return &e
}

func (e *WebEngine) Serve(w http.ResponseWriter, r *http.Request) {
	e.log.Debug("Serve")
	tmpl := template.Must(template.ParseFiles("cmd/web/ui/index.html"))
	tmpl.Execute(w, nil)
}
