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

	// get genres, if any
	genres, err := m.getGenres(ctx, song)
	if err != nil {
		return nil, err
	}

	song.SongGenre = genres

	return &song, nil
}

// All returns all songs and error, if any
func (m *DBModel) All() ([]*Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, rating, duration, riaa_rating, 
	created_at, updated_at from songs order by title`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*Song

	for rows.Next() {
		var song Song
		err := rows.Scan(
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

		// get genres, if any
		genres, err := m.getGenres(ctx, song)
		if err != nil {
			return nil, err
		}

		song.SongGenre = genres
		songs = append(songs, &song)
	}

	return songs, nil
}

func (m *DBModel) getGenres(ctx context.Context, song Song) (map[int]string, error) {
	genreQuery := `select 
				sg.id, sg.song_id, sg.genre_id, g.genre_name
			from
				songs_genres sg
				left join genres g on (g.id = sg.genre_id)
			where
				sg.song_id = $1`
	genreRows, _ := m.DB.QueryContext(ctx, genreQuery, song.ID)

	genres := make(map[int]string)
	for genreRows.Next() {
		var sg SongGenre
		err := genreRows.Scan(
			&sg.ID,
			&sg.SongID,
			&sg.GenreID,
			&sg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[sg.ID] = sg.Genre.GenreName
	}
	genreRows.Close()

	return genres, nil

}
