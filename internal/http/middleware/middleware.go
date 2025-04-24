package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func NewChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			next = middleware(next)

		}
		return next
	}
}

func DefaultChain(router *http.ServeMux) http.Handler {
	chain := NewChain(
		LogRequest,
	)
	return chain(router)
}
