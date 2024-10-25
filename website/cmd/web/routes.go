package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", app.home)

	router.PathPrefix("/static/").Handler(app.static("./ui/static/"))
	return router
}
