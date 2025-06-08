package controller

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/keribend/hello-web/internal/service"
	"github.com/keribend/hello-web/internal/views/components"
)

type HTMLController struct {
	service *service.Service
}

func NewHTMLController(s *service.Service) *HTMLController {
	return &HTMLController{s}
}

func (s *HTMLController) ListEvents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	events, err := s.service.FindAllEvents(ctx)
	if err != nil {
		log.Println("FindAllEvents error: ", err)
		return
	}

	log.Println("events: ", events)

	s.render(w, r, http.StatusOK, components.EventList(events))
}

func (s *HTMLController) AddCheckinToEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	eventId, err := ParseId64(r.PathValue("id"))
	if err != nil {
		log.Println("strconv.Atoi error: ", err)
	}

	err = s.service.AddCheckinToEvent(ctx, eventId)
	if err != nil {
		log.Println("AddCheckinToEvent error: ", err)
	}

	s.renderAlert(w, r, "Checkin done!")
}

func (s *HTMLController) Checkins(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (s *HTMLController) render(w http.ResponseWriter, r *http.Request, com templ.Component) {
	if err := com.Render(r.Context(), w); err != nil {
		log.Println("Render error: ", err)
	}
}

func (s *HTMLController) renderAlert(w http.ResponseWriter, r *http.Request, msg string) {
	s.hxRetarget(w, "body")
	s.hxReswap(w, "beforeend")
	s.render(w, r, components.AlertSuccess(msg))
}

func (s *HTMLController) hxRetarget(w http.ResponseWriter, target string) {
	w.Header().Set("HX-Retarget", target)
}

func (s *HTMLController) hxReswap(w http.ResponseWriter, swap string) {
	w.Header().Set("HX-Reswap", swap)
}
