package main

import (
	"bedandbreakfast/pkg/config"
	"bedandbreakfast/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurve)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}