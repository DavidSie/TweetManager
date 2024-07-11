package handlers

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/DavidSie/TweetManager/internal/config"
	"github.com/DavidSie/TweetManager/internal/helpers"
	"github.com/go-chi/chi"
	"github.com/go-chi/httprate"
)

var app config.AppConfig

func TestMain(m *testing.M) {

	// change this to true when in production
	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	repo := NewTestRepo(&app)
	NewHandlers(repo)
	helpers.NewHelpers(&app)

	os.Exit(m.Run())
}

func getRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(httprate.LimitAll(1000, 1*time.Minute))
	mux.Use(httprate.LimitByIP(100, 1*time.Minute))

	mux.Get("/", http.HandlerFunc(Repo.Home))
	mux.Get("/tweets", http.HandlerFunc(Repo.TweetsJSON))
	mux.Get("/tweets-with-emotions", http.HandlerFunc(Repo.TweetsWithEmotionsJSON))

	return mux
}
