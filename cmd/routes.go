package main

import (
	"net/http"
	"time"

	"github.com/DavidSie/TweetManager/internal/config"
	"github.com/DavidSie/TweetManager/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(httprate.LimitAll(1000, 1*time.Minute))
	mux.Use(httprate.LimitByIP(100, 1*time.Minute))

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/tweets", http.HandlerFunc(handlers.Repo.TweetsJSON))
	mux.Get("/tweets-with-emotions", http.HandlerFunc(handlers.Repo.TweetsWithEmotionsJSON))

	return mux
}
