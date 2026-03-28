package router

import (
	"net/http"
	"onlinekino/backend/handlers"
)

func SetupRouter() {
	http.HandleFunc("/", handlers.MainPage)

	fs := http.FileServer(http.Dir("./frontend/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/watch/film/", handlers.WatchFilmPage)

	http.HandleFunc("/health", handlers.HealthHandler)
}
