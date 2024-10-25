package main

import "github.com/gorilla/mux"

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", app.all).Methods("GET")

	return router
}
