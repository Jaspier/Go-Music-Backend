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
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Artist      string         `json:"artist"`
	Year        int            `json:"year"`
	ReleaseDate time.Time      `json:"release_date"`
	Duration    int            `json:"duration"`
	Rating      int            `json:"rating"`
	RIAARating  string         `json:"riaa_rating"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	SongGenre   map[int]string `json:"genres"`
	Cover       string         `json:"cover"`
}

// Genre is the type for genre
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// SongGenre is the type for song genre
type SongGenre struct {
	ID        int       `json:"-"`
	SongID    int       `json:"-"`
	GenreID   int       `json:"-"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// User is the type for users
type User struct {
	ID       int
	Email    string
	Password string
}
