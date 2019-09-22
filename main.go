package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ianadiwibowo/scout-regiment/app"
	"github.com/ianadiwibowo/scout-regiment/db"
	"github.com/subosito/gotenv"
)

// Starting point
func main() {
	loadEnvironmentVariables()

	router := mux.NewRouter().StrictSlash(true)
	db := db.SetupDatabase()

	app := &app.App{
		Router: router,
		DB:     db,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

// Load system configurations, get them via: os.Getenv("{ENV_KEY_NAME}")
func loadEnvironmentVariables() {
	if os.Getenv("APP_ENV") != "production" {
		err := gotenv.Load()
		if err != nil {
			panic(err)
		}
	}
}
