package mainhandler

import "net/http"

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/main.html")
	}
}
