package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneSong(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	song, err := app.models.DB.Get(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, song, "song")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, songs, "songs")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.GenresAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// func (app *application) deleteSong(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) insertSong(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) updateSong(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) searchSongs(w http.ResponseWriter, r *http.Request) {

// }
