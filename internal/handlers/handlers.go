package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/DavidSie/TweetManager/internal/config"
	"github.com/DavidSie/TweetManager/internal/driver"
	"github.com/DavidSie/TweetManager/internal/helpers"
	"github.com/DavidSie/TweetManager/internal/repository"
	"github.com/DavidSie/TweetManager/internal/repository/dbrepo"
	"github.com/DavidSie/TweetManager/pkg/model"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPotgresRepo(db.SQL, a),
	}
}

// NewReposito creates a new repository for testing purposes
func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewTestingRepo(a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	welcomeMsg := "Welcome to Tweeter Manager"

	_, err := w.Write([]byte(welcomeMsg))
	if err != nil {
		m.App.ErrorLog.Println(err)
	}
	//render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

// TweetsJSON is a function that uses parameters: symbol, start_date, end_date  from query database for the tweets in this time frame
//
// curl "http://localhost:8080/tweets?symbol=symbol&start_date=2024-06-24&end_date=2024-07-09"
func (m *Repository) TweetsJSON(w http.ResponseWriter, r *http.Request) {
	resp := model.TweetResponse{}

	symbol := r.URL.Query().Get("symbol")
	if len(symbol) == 0 {
		helpers.ClientError(w, http.StatusBadRequest, "symbol not set as query parameter")
		return
	}
	sd := r.URL.Query().Get("start_date")

	if len(sd) == 0 {
		helpers.ClientError(w, http.StatusBadRequest, "start_date not set as query parameter")
		return
	}

	ed := r.URL.Query().Get("end_date")
	if len(ed) == 0 {
		helpers.ClientError(w, http.StatusBadRequest, "end_date not set as query parameter")
		return
	}

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w, errors.New("error can't parse start date: "+err.Error()))
		return
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, errors.New("error can't parse end date: "+err.Error()))
		return
	}

	tweets, err := m.DB.GetTweetsBySymbolByDate(symbol, startDate, endDate)
	if err != nil {
		helpers.ServerError(w, errors.New("error can't get tweets by the symbol by date: "+err.Error()))
		return
	}

	resp.Tweets = tweets
	resp.OK = true
	out, err := json.MarshalIndent(resp, "", "	")
	if err != nil {
		helpers.ServerError(w, errors.New("error can't marshal response into json: "+err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		m.App.ErrorLog.Println(err)
	}
}

// Home is the home page handler
func (m *Repository) TweetsWithEmotionsJSON(w http.ResponseWriter, r *http.Request) {

	resp := model.TweetResponse{}
	out, err := json.MarshalIndent(resp, "", "	")
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		m.App.ErrorLog.Println(err)
	}
	//render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}
