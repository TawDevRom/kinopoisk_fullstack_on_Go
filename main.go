package main

import (
	"log"
	"net/http"
	"onlinekino/backend/config"
	"onlinekino/backend/router"
)

func main() {
	cfg := config.LoadConfig()

	router.SetupRouter()

	log.Println("Сервер запущен на :" + cfg.Port)

	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// ...:8080/film/film_id/watch
// ...:8080/serial/serial_id/watch?episode_number=&season_number=5

//http://localhosh:8080/watch/film/1
