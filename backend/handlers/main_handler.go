package handlers

import "net/http"

func MainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "backend/static/index.html")
}
