package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App is the core system of this application
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// SetupRouter initializes the API endpoint routes
func (app *App) SetupRouter() {
	app.Router.Methods("GET").Path("/cats/{id}").HandlerFunc(app.getCatHandler)
	app.Router.Methods("POST").Path("/cats").HandlerFunc(app.createCatHandler)
}

// GET /cats/{id}
func (app *App) getCatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You called GET /cats/{id}")
}

// POST /cats
func (app *App) createCatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You called POST /cats")
}
