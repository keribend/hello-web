package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/keribend/hello-web/internal/service"
)

type APIController struct {
	service *service.Service
}

func NewAPIController(s *service.Service) *APIController {
	return &APIController{s}
}

func (c *APIController) ListEvents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	events, err := c.service.FindAllEvents(ctx)
	if err != nil {
		log.Println("FindAllEvents error: ", err)
		return
	}

	log.Println("events: ", events)

	c.response(w, http.StatusOK, events)
}

func (c *APIController) GetEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := ParseId64(r.PathValue("id"))
	if err != nil {
		c.errorResponse(w, err)
		return
	}

	event, err := c.service.FindEvent(ctx, id)
	if err != nil {
		c.errorResponse(w, err)
		return
	}

	c.successResponse(w, event)
}

func (c *APIController) response(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("json.NewEncoder().Encode() error: ", err)
	}
}

func (c *APIController) errorResponse(w http.ResponseWriter, err error) {
	c.response(w, http.StatusInternalServerError, err)
}

func (c *APIController) successResponse(w http.ResponseWriter, data interface{}) {
	c.response(w, http.StatusOK, data)
}
