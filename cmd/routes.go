package main

import (
	"net/http"

	"github.com/DavidSie/TweetManager/internal/config"
	"github.com/DavidSie/TweetManager/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// mux.Use(middleware.Recoverer)
	// mux.Use(NoSurv)
	// mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/tweets", http.HandlerFunc(handlers.Repo.TweetsJSON))
	mux.Get("/tweets-with-emotions", http.HandlerFunc(handlers.Repo.TweetsWithEmotionsJSON))

	// fileServer := http.FileServer(http.Dir("./static/"))
	// mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
