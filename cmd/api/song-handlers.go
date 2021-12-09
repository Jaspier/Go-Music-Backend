package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type jsonResp struct {
	OK      bool   `json: "ok"`
	Message string `json:"message"`
}

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

func (app *application) getAllSongsByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	songs, err := app.models.DB.All(genreID)
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

type SongPayload struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Year        string `json:"year"`
	ReleaseDate string `json:"release_date"`
	Duration    string `json:"duration"`
	Rating      string `json:"rating"`
	RIAARating  string `json:"riaa_rating"`
}

func (app *application) editSong(w http.ResponseWriter, r *http.Request) {
	var payload SongPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var song models.Song

	if payload.ID != "0" {
		id, _ := strconv.Atoi(payload.ID)
		s, _ := app.models.DB.Get(id)
		song = *s
		song.UpdatedAt = time.Now()
	}

	song.ID, _ = strconv.Atoi(payload.ID)
	song.Title = payload.Title
	song.Artist = payload.Artist
	song.ReleaseDate, _ = time.Parse("2006-01-02", payload.ReleaseDate)
	song.Year = song.ReleaseDate.Year()
	song.Duration, _ = strconv.Atoi(payload.Duration)
	song.Rating, _ = strconv.Atoi(payload.Rating)
	song.RIAARating = payload.RIAARating
	song.CreatedAt = time.Now()
	song.UpdatedAt = time.Now()

	if song.ID == 0 {
		err = app.models.DB.InsertSong(song)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		err = app.models.DB.UpdateSong(song)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

// func (app *application) deleteSong(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) updateSong(w http.ResponseWriter, r *http.Request) {

// }

// func (app *application) searchSongs(w http.ResponseWriter, r *http.Request) {

// }
