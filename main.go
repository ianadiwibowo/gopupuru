package main

import (
	"net/http"
	"os"

	"github.com/ianadiwibowo/gopupuru/app"
	"github.com/subosito/gotenv"
)

// Starting point
func main() {
	loadEnvironmentVariables()

	app := app.New()
	app.SetupRouter()

	http.ListenAndServe(":8080", app.Router)
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
