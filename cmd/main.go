package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DavidSie/TweetManager/internal/config"
	"github.com/spf13/viper"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	log.Println("Tweet Manager Starts")
	defer log.Println("Tweet Manager Stops")

	err := run()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Starting application on port %s \n", portNumber)
	srv := &http.Server{Addr: portNumber, Handler: routes(&app)}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	viper.SetConfigName("secrets")              // name of config file (without extension)
	viper.SetConfigType("yaml")                 // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/tweet-manager/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.tweet-manager") // call multiple times to add many search paths
	viper.AddConfigPath(".")                    // optionally look for config in the working directory
	err := viper.ReadInConfig()                 // Find and read the config file
	if err != nil {                             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error secrets file: %w", err))
	}

	// change this to true when in production
	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	return nil
}
