package main

import (
	"gopkg.in/pg.v4"
)

// Movie all values
type Movie struct {
	ID          int64    `sql:"id"`
	Section     string   `sql:"section"`
	Name        string   `sql:"name"`
	EngName     string   `sql:"eng_name"`
	Year        int64    `sql:"year"`
	Genre       []string `sql:"genre"        pg:",array" `
	Country     []string `sql:"country"      pg:",array"`
	RawCountry  string   `sql:"raw_country"`
	Director    []string `sql:"director"     pg:",array"`
	Producer    []string `sql:"producer"     pg:",array"`
	Actor       []string `sql:"actor"        pg:",array"`
	Description string   `sql:"description"`
	Age         string   `sql:"age"`
	ReleaseDate string   `sql:"release_date"`
	RussianDate string   `sql:"russian_date"`
	Duration    string   `sql:"duration"`
	Kinopoisk   float64  `sql:"kinopoisk"`
	IMDb        float64  `sql:"imdb"`
	Poster      string   `sql:"poster"`
	PosterURL   string   `sql:"poster_url"`
	Torrent     []Torrent `sql:"-"`
}

// Torrent all values
type Torrent struct {
	ID         int64   `sql:"id"`
	MovieID    int64   `sql:"movie_id"`
	DateCreate string  `sql:"date_create"`
	Href       string  `sql:"href"`
	Torrent    string  `sql:"torrent"`
	Magnet     string  `sql:"magnet"`
	NNM        float64 `sql:"nnm"`
	// SubtitlesType string  `sql:"subtitles_type"`
	// Subtitles     string  `sql:"subtitles"`
	Video      string `sql:"video"`
	Quality    string `sql:"quality"`
	Resolution string `sql:"resolution"`
	// Audio1        string  `sql:"audio1"`
	// Audio2        string  `sql:"audio2"`
	// Audio3        string  `sql:"audio3"`
	Translation string `sql:"translation"`
	Size        int64  `sql:"size"`
	Seeders     int64  `sql:"seeders"`
	Leechers    int64  `sql:"leechers"`
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

func (app *application) getMovie() Movie {
	var movie Movie
	app.database.Model(&movie).First()
	return movie
}

func (app *application) getMovies() []Movie {
	var movies []Movie
	app.database.Model(&movies).Select()
	for i, movie := range movies {
		movies[i].Torrent = app.getMovieTorrents(movie.ID)
	}
	return movies
}

func (app *application) getMovieTorrents(id int64) []Torrent {
	var torrents []Torrent
	app.database.Model(&torrents).Where("movie_id = ?", id).Select()
	return torrents
}
