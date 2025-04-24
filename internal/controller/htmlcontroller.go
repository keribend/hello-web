package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/keribend/hello-web/internal/service"
	"github.com/keribend/hello-web/internal/views/components"
)

type HtmlController struct {
	service *service.Service
}

func New(s *service.Service) *HtmlController {
	return &HtmlController{s}
}

func (s *HtmlController) EventList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	events, err := s.service.FindAllEvents(ctx)
	if err != nil {
		log.Println("FindAllEvents error: ", err)
		return
	}

	log.Println("events: ", events)

	s.render(w, r, http.StatusOK, components.EventList(events))
}

func (s *HtmlController) AddCheckinToEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// eventId := r.URL.
	/// log.Println("AddCheckinToEvent triggered for eventId: ", chi.URLParam(r, "eventId"))

	eventId, err := strconv.ParseInt(chi.URLParam(r, "eventId"), 10, 64)
	if err != nil {
		log.Println("strconv.Atoi error: ", err)
	}

	err = s.service.AddCheckinToEvent(ctx, eventId)
	if err != nil {
		log.Println("AddCheckinToEvent error: ", err)
	}

	s.renderAlert(w, r, http.StatusOK, "Checkin done!")
}

func (s *HtmlController) Checkins(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (s *HtmlController) render(w http.ResponseWriter, r *http.Request, status int, com templ.Component) {
	if err := com.Render(r.Context(), w); err != nil {
		log.Println("Render error: ", err)
	}
	w.WriteHeader(status)
}

func (s *HtmlController) renderAlert(w http.ResponseWriter, r *http.Request, status int, msg string) {
	s.hxRetarget(w, "body")
	s.hxReswap(w, "beforeend")
	s.render(w, r, status, components.AlertSuccess(msg))
}

func (s *HtmlController) hxRetarget(w http.ResponseWriter, target string) {
	w.Header().Set("HX-Retarget", target)
}

func (s *HtmlController) hxReswap(w http.ResponseWriter, swap string) {
	w.Header().Set("HX-Reswap", swap)
}
