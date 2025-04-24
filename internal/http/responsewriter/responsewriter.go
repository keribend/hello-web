package responsewriter

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func New(w http.ResponseWriter, statusConde int) ResponseWriter {
	return ResponseWriter{
		ResponseWriter: w,
		StatusCode:     statusConde,
	}
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}
