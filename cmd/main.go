package main

import (
	"log"
	"net/http"

	mainhandler "github.com/SergeyGolang/portfolio-go/internal/handlers/mainHandler"
	"github.com/go-chi/chi"
)

func main() {
	//TODO: init logger

	//TODO: middleware

	//TODO: init router
	router := chi.NewRouter()

	router.Get("/", mainhandler.New())
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//TODO: start server
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatal("failed to start server")
	}
}
