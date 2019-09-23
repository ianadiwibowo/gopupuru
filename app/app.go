package app

import (
	"os"

	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"github.com/ianadiwibowo/scout-regiment/db"
	"github.com/ianadiwibowo/scout-regiment/delivery/handler"
	"github.com/ianadiwibowo/scout-regiment/delivery/serializer"
	"github.com/jinzhu/gorm"
)

// App is the core system of this application
type App struct {
	Router     *mux.Router
	DB         *gorm.DB
	Serializer *serializer.Serializer
}

// New initializes the app
func New() *App {
	return &App{
		Router:     mux.NewRouter().StrictSlash(true),
		DB:         db.New(),
		Serializer: serializer.New(),
	}
}

// SetupRouter initializes the API endpoint routes
func (app *App) SetupRouter() {
	// Authenticate using basic auth
	app.Router.Use(httpauth.SimpleBasicAuth(os.Getenv("BASIC_USERNAME"), os.Getenv("BASIC_PASSWORD")))

	catsHandler := handler.NewCatsHandler(app.DB, app.Serializer)
	// app.Router.Methods("GET").Path("/cats").HandlerFunc(catsHandler.Index)
	app.Router.Methods("GET").Path("/cats/{id}").HandlerFunc(catsHandler.Show)
	app.Router.Methods("POST").Path("/cats").HandlerFunc(catsHandler.Create)
	// app.Router.Methods("PATCH").Path("/cats/{id}").HandlerFunc(catsHandler.Update)
	// app.Router.Methods("DELETE").Path("/cats/{id}").HandlerFunc(catsHandler.Delete)
}
