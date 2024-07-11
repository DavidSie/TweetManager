package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/DavidSie/TweetManager/internal/config"
)

var app *config.AppConfig

// Sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int, extraText string) {
	app.InfoLog.Println("Client error with status of", status)
	errorMsg := http.StatusText(status)
	if extraText != "" {
		errorMsg += "\n" + extraText
	}
	http.Error(w, errorMsg, status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
