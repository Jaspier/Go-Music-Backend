package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/song/:id", app.getOneSong)
	router.HandlerFunc(http.MethodGet, "/v1/songs", app.getAllSongs)
	router.HandlerFunc(http.MethodGet, "/v1/songs/:genre_id", app.getAllSongsByGenre)

	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)

	router.HandlerFunc(http.MethodPost, "/v1/admin/editsong", app.editSong)

	return app.enableCORS(router)
}
