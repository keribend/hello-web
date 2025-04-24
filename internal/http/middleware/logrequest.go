package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/keribend/hello-web/internal/http/responsewriter"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		resp := responsewriter.New(w, http.StatusOK)
		next.ServeHTTP(&resp, r)
		log.Println(resp.StatusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
