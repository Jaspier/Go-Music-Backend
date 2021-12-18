package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()

	secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)

	router.HandlerFunc(http.MethodGet, "/v1/song/:id", app.getOneSong)
	router.HandlerFunc(http.MethodGet, "/v1/songs", app.getAllSongs)
	router.HandlerFunc(http.MethodGet, "/v1/songs/:genre_id", app.getAllSongsByGenre)

	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)

	router.POST("/v1/admin/editsong", app.wrap(secure.ThenFunc(app.editSong)))

	router.GET("/v1/admin/deletesong/:id", app.wrap(secure.ThenFunc(app.deleteSong)))

	return app.enableCORS(router)
}
