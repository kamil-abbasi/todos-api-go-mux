package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestMethod := r.Method
		requestUrl := r.URL

		log.Printf("%v %v", requestMethod, requestUrl)

		next.ServeHTTP(w, r)
	})
}
