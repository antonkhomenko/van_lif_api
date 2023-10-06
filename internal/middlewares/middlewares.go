package middlewares

import (
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func sendJson(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendJson(&w)
		enableCors(&w)
		log.Printf("user %s request %s\n", r.RemoteAddr, r.URL)
		next(w, r)
	}
}
