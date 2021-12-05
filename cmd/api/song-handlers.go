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

	app.logger.Println("id is", id)

	song, err := app.models.DB.Get(id)

	// song := models.Song{
	// 	ID:          id,
	// 	Title:       "Some song",
	// 	Description: "Some description",
	// 	Year:        2021,
	// 	ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
	// 	Duration:    230,
	// 	Rating:      5,
	// 	RIAARating:  "PG-13",
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// }

	err = app.writeJSON(w, http.StatusOK, song, "song")
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getAllSongs(w http.ResponseWriter, r *http.Request) {

}
