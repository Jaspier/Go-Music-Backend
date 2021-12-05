package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Get returns one song and error, if any
func (m *DBModel) Get(id int) (*Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, rating, duration, riaa_rating, 
				created_at, updated_at from songs where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var song Song

	err := row.Scan(
		&song.ID,
		&song.Title,
		&song.Description,
		&song.Year,
		&song.ReleaseDate,
		&song.Rating,
		&song.Duration,
		&song.RIAARating,
		&song.CreatedAt,
		&song.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &song, nil
}

// All returns all songs and error, if any
func (m *DBModel) All(id int) ([]*Song, error) {
	return nil, nil
}
