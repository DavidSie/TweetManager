package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DavidSie/TweetManager/internal/config"
	"github.com/DavidSie/TweetManager/pkg/model"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	// DB  repository.DatabaseRepo
}

func NewRepo(a *config.AppConfig) *Repository { //, db *driver.DB
	return &Repository{
		App: a,
		// DB:  dbrepo.NewPotgresRepo(db.SQL, a),
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

// Home is the home page handler
func (m *Repository) TweetsJSON(w http.ResponseWriter, r *http.Request) {

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
