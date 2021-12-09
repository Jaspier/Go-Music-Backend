package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Get returns one song and error, if any
func (m *DBModel) Get(id int) (*Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, artist, year, release_date, rating, duration, riaa_rating, 
				created_at, updated_at from songs where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var song Song

	err := row.Scan(
		&song.ID,
		&song.Title,
		&song.Artist,
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
func (m *DBModel) All(genre ...int) ([]*Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(genre) > 0 {
		where = fmt.Sprintf("where id in (select song_id from songs_genres where genre_id = %d)", genre[0])
	}

	query := fmt.Sprintf(`select id, title, artist, year, release_date, rating, duration, riaa_rating, 
	created_at, updated_at from songs %s order by title`, where)

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
			&song.Artist,
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

func (m *DBModel) GenresAll() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, genre_name, created_at, updated_at from genres order by genre_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*Genre

	for rows.Next() {
		var g Genre
		err := rows.Scan(
			&g.ID,
			&g.GenreName,
			&g.CreatedAt,
			&g.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &g)
	}

	return genres, nil
}

func (m *DBModel) InsertSong(song Song) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into songs (title, artist, year, release_date, duration, rating, riaa_rating,
				created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.ExecContext(ctx, stmt,
		song.Title,
		song.Artist,
		song.Year,
		song.ReleaseDate,
		song.Duration,
		song.Rating,
		song.RIAARating,
		song.CreatedAt,
		song.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) UpdateSong(song Song) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update songs set title = $1, artist = $2, year = $3, release_date = $4, 
				duration = $5, rating = $6, riaa_rating = $7,
				updated_at = $8 where id = $9`

	_, err := m.DB.ExecContext(ctx, stmt,
		song.Title,
		song.Artist,
		song.Year,
		song.ReleaseDate,
		song.Duration,
		song.Rating,
		song.RIAARating,
		song.UpdatedAt,
		song.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) DeleteSong(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "delete from songs where id = $1"

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
