package handlers

import "net/http"

func WatchFilmPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "backend/static/player.html")
}
