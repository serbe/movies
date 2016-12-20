package main

import (
	"database/sql"

	"github.com/lib/pq"
)

func scanMovie(r *sql.Row) (Movie, error) {
	var s Movie
	if err := r.Scan(
		&s.ID,
		&s.Section,
		&s.Name,
		&s.EngName,
		&s.Year,
		pq.Array(&s.Genre),
		pq.Array(&s.Country),
		&s.RawCountry,
		pq.Array(&s.Director),
		pq.Array(&s.Producer),
		pq.Array(&s.Actor),
		&s.Description,
		&s.Age,
		&s.ReleaseDate,
		&s.RussianDate,
		&s.Duration,
		&s.Kinopoisk,
		&s.IMDb,
		&s.Poster,
		&s.PosterURL,
		// &s.CreatedAt,
		// &s.UpdatedAt,
		// &s.Torrent,
		// &s.NNM,
	); err != nil {
		return Movie{}, err
	}
	return s, nil
}

func scanMovies(rs *sql.Rows) ([]Movie, error) {
	structs := make([]Movie, 0, 16)
	var err error
	for rs.Next() {
		var s Movie
		if err = rs.Scan(
			&s.ID,
			&s.Section,
			&s.Name,
			&s.EngName,
			&s.Year,
			&s.Genre,
			&s.Country,
			&s.RawCountry,
			&s.Director,
			&s.Producer,
			&s.Actor,
			&s.Description,
			&s.Age,
			&s.ReleaseDate,
			&s.RussianDate,
			&s.Duration,
			&s.Kinopoisk,
			&s.IMDb,
			&s.Poster,
			&s.PosterURL,
			&s.CreatedAt,
			&s.UpdatedAt,
			// &s.Torrent,
			// &s.NNM,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func scanTorrent(r *sql.Row) (Torrent, error) {
	var s Torrent
	if err := r.Scan(
		&s.ID,
		&s.MovieID,
		&s.DateCreate,
		&s.Href,
		&s.Torrent,
		&s.Magnet,
		&s.NNM,
		&s.SubtitlesType,
		&s.Subtitles,
		&s.Video,
		&s.Quality,
		&s.Resolution,
		&s.Audio1,
		&s.Audio2,
		&s.Audio3,
		&s.Translation,
		&s.Size,
		&s.Seeders,
		&s.Leechers,
		&s.CreatedAt,
		&s.UpdatedAt,
	); err != nil {
		return Torrent{}, err
	}
	return s, nil
}

func scanTorrents(rs *sql.Rows) ([]Torrent, error) {
	structs := make([]Torrent, 0, 16)
	var err error
	for rs.Next() {
		var s Torrent
		if err = rs.Scan(
			&s.ID,
			&s.MovieID,
			&s.DateCreate,
			&s.Href,
			&s.Torrent,
			&s.Magnet,
			&s.NNM,
			&s.SubtitlesType,
			&s.Subtitles,
			&s.Video,
			&s.Quality,
			&s.Resolution,
			&s.Audio1,
			&s.Audio2,
			&s.Audio3,
			&s.Translation,
			&s.Size,
			&s.Seeders,
			&s.Leechers,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func scanSearch(r *sql.Row) (search, error) {
	var s search
	if err := r.Scan(
		&s.ID,
		&s.MovieID,
	); err != nil {
		return search{}, err
	}
	return s, nil
}

func scanSearchs(rs *sql.Rows) ([]search, error) {
	structs := make([]search, 0, 16)
	var err error
	for rs.Next() {
		var s search
		if err = rs.Scan(
			&s.ID,
			&s.MovieID,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}
