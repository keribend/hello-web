package routing

import (
	"net/http"

	"github.com/keribend/hello-web/internal/controller"
)

func NewAPIRouter(c *controller.APIController) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /events", c.ListEvents)
	r.HandleFunc("GET /events/{id}", c.GetEvent)

	return r
}
