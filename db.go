package main

import (
	"gopkg.in/pg.v4"
)

// Movie all values
type Movie struct {
	ID          int64
	Name        string
	EngName     string
	Year        int64
	Genre       []string
	Country     string
	Director    []string
	Actors      []string
	Description string
	ReleaseDate string
	RussianDate string
	Duration    string
	Kinopoisk   float64
	IMDb        float64
	Poster      string
	PosterURL   string
	NNM         float64
	Torrents    []Torrent
	Len         int
}

// Torrent all values
type Torrent struct {
	ID         int64
	MovieID    int64
	DateCreate string
	Href       string
	// Torrent       string
	Magnet string
	NNM    float64
	// SubtitlesType string
	// Subtitles     string
	Video      string
	Quality    string
	Resolution string
	// Audio1        string
	// Audio2        string
	// Audio3        string
	Translation string
	Size        int64
	Seeders     int64
	Leechers    int64
}

func (app *application) initDB() {
	db := pg.Connect(&pg.Options{
		Database: app.config.Base.Dbname,
		User:     app.config.Base.User,
		Password: app.config.Base.Password,
		SSL:      app.config.Base.Sslmode,
	})
	app.database = db
}

func (app *application) getMovie() (Movie) {
	var movie Movie
	app.database.Model(&movie).Limit(1).Select()
	return movie
}

func (app *application) getMovies() ([]Movie) {
	var movies []Movie
	app.database.Model(&movies).Select()
	return movies
}