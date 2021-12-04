package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/song/:id", app.getOneSong)
	router.HandlerFunc(http.MethodGet, "/v1/songs", app.getAllSongs)

	return router
}
