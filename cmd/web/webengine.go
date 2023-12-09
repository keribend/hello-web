package main

import (
	"html/template"
	"net/http"

	"go.uber.org/zap"
)

type Counter struct {
	ID    int
	Name  string
	Value uint
}

var counters = []Counter{}

type WebEngine struct {
	log       *zap.Logger
	templates *template.Template
}

func NewWebEngine() *WebEngine {
	e := WebEngine{}
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	e.log = log
	e.templates, err = e.ParseTemplates()
	if err != nil {
		e.log.Error("ParseTemplates", zap.Error(err))
	}
	return &e
}

func (e *WebEngine) ParseTemplates() (*template.Template, error) {
	return template.ParseGlob("cmd/web/ui/*")

}

func (e *WebEngine) Serve(w http.ResponseWriter, r *http.Request) {
	e.log.Debug("Serve")
	e.templates.ExecuteTemplate(w, "homeHTML", counters)
}

func (e *WebEngine) GetCounters(w http.ResponseWriter, r *http.Request) {
	e.log.Debug("GetCounters")
	e.templates.ExecuteTemplate(w, "countersHTML", counters)
}

func (e *WebEngine) CreateCounter(w http.ResponseWriter, r *http.Request) {
	e.log.Debug("CreateCounter")
}

func (e *WebEngine) IncreaseCounter(w http.ResponseWriter, r *http.Request) {
	e.log.Debug("IncreaseCounter")
}

func (e *WebEngine) DecreaseCounter(w http.ResponseWriter, r *http.Request) {
	e.log.Debug("DecreaseCounter")
}
