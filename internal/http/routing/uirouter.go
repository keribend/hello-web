package routing

import (
	"net/http"

	"github.com/keribend/hello-web/internal/controller"
)

func NewUIRouter(c *controller.HTMLController) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", c.ListEvents)
	r.HandleFunc("POST /events/{id}/checkin", c.AddCheckinToEvent)

	// r.Post("/events/{eventId}/checkin", controller.AddCheckinToEvent)
	// r.Get("/checkins", controller.Checkins)

	return r
}
