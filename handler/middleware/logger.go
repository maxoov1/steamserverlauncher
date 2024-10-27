package middleware

import (
	"log"
	"net/http"
	"time"
)

type statusCodeResponseWriter struct {
	responseWriter http.ResponseWriter
	statusCode     int
}

func NewStatusCodeResponseWriter(responseWriter http.ResponseWriter) *statusCodeResponseWriter {
	return &statusCodeResponseWriter{responseWriter: responseWriter, statusCode: http.StatusOK}
}

func (w *statusCodeResponseWriter) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *statusCodeResponseWriter) Write(buffer []byte) (int, error) {
	return w.responseWriter.Write(buffer)
}

func (w *statusCodeResponseWriter) WriteHeader(statusCode int) {
	if statusCode != 0 {
		w.statusCode = statusCode
	}

	w.responseWriter.WriteHeader(statusCode)
}

func (w *statusCodeResponseWriter) StatusCode() int {
	return w.statusCode
}

type Logger struct {
	handler http.Handler
}

func NewLogger(handler http.Handler) *Logger {
	return &Logger{handler: handler}
}

func (m *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	responseWriter := NewStatusCodeResponseWriter(w)
	m.handler.ServeHTTP(responseWriter, r)
	log.Printf("%d %s %s %v", responseWriter.StatusCode(), r.Method, r.URL.Path, time.Since(start))
}
