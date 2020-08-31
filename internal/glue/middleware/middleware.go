package middleware

import (
	"net/http"
)

func ContentTypeJson(handler http.Handler) http.Handler {
	return ContentType(handler, "application/json")
}

func ContentType(handler http.Handler, contentType string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", contentType)
		handler.ServeHTTP(rw, r)
	})
}
