package application

import (
	"fmt"
	"net/http"
)

func customHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Expert-Systems", "Miyuki")
		w.Header().Add("Server", "net/http")
		next.ServeHTTP(w, r)
	})
}

func httpsRedirect(applicationMode string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if applicationMode == "production" {
				if r.Header.Get("X-Forwarded-Proto") != "https" {
					sslUrl := fmt.Sprintf("https://%s%s", r.Host, r.RequestURI)
					http.Redirect(w, r, sslUrl, http.StatusPermanentRedirect)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
