package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Song is the type for songs
type Song struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Year        int         `json:"year"`
	ReleaseDate time.Time   `json:"release_date"`
	Duration    int         `json:"duration"`
	Rating      int         `json:"rating"`
	RIAARating  string      `json:"riaa_rating"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	SongGenre   []SongGenre `json:"-"`
}

// Genre is the type for genre
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SongGenre is the type for song genre
type SongGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movie_id"`
	GenreID   int       `json:"genre_id"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
