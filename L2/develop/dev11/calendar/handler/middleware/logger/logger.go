package logger

import (
	"log"
	"net/http"
	"os"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Wrap(handler http.Handler) http.Handler {
	logger := log.New(os.Stdout, "LOG ", log.Ldate|log.Ltime)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     200,
		}

		start := time.Now()
		handler.ServeHTTP(rw, r)
		duration := time.Since(start)

		logger.Printf("method=%s path=%s duration=%s status=%d", r.Method, r.RequestURI, duration, rw.statusCode)
	})
}
