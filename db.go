package main

import (
	"gopkg.in/pg.v4"
)

// Data - data to sent JSON
type Data struct {
	Offset int
	Count  int
	Limit  int
	ImgDir string
	Movies []Movie
}

// Movie all values
type Movie struct {
	ID          int       `sql:"id"`
	Section     string    `sql:"section"`
	Name        string    `sql:"name"`
	EngName     string    `sql:"eng_name"`
	Year        int       `sql:"year"`
	Genre       []string  `sql:"genre"        pg:",array" `
	Country     []string  `sql:"country"      pg:",array"`
	RawCountry  string    `sql:"raw_country"`
	Director    []string  `sql:"director"     pg:",array"`
	Producer    []string  `sql:"producer"     pg:",array"`
	Actor       []string  `sql:"actor"        pg:",array"`
	Description string    `sql:"description"`
	Age         string    `sql:"age"`
	ReleaseDate string    `sql:"release_date"`
	RussianDate string    `sql:"russian_date"`
	Duration    string    `sql:"duration"`
	Kinopoisk   float64   `sql:"kinopoisk"`
	IMDb        float64   `sql:"imdb"`
	Poster      string    `sql:"poster"`
	PosterURL   string    `sql:"poster_url"`
	Torrent     []Torrent `sql:"-"`
	NNM         float64   `sql:"-"`
}

// Torrent all values
type Torrent struct {
	ID          int     `sql:"id"`
	MovieID     int     `sql:"movie_id"`
	DateCreate  string  `sql:"date_create"`
	Href        string  `sql:"href"`
	Torrent     string  `sql:"torrent"`
	Magnet      string  `sql:"magnet"`
	NNM         float64 `sql:"nnm"`
	Video       string  `sql:"video"`
	Quality     string  `sql:"quality"`
	Resolution  string  `sql:"resolution"`
	Translation string  `sql:"translation"`
	Size        int     `sql:"size"`
	Seeders     int     `sql:"seeders"`
	Leechers    int     `sql:"leechers"`
	// SubtitlesType string  `sql:"subtitles_type"`
	// Subtitles     string  `sql:"subtitles"`
	// Audio1        string  `sql:"audio1"`
	// Audio2        string  `sql:"audio2"`
	// Audio3        string  `sql:"audio3"`
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

func (app *application) getMovies(limit int, offset int) Data {
	var (
		movies, m []Movie
		count     int
		data      Data
	)

	count, _ = app.database.Model(&movies).Count()
	if limit == 0 {
		limit = count
	}
	if offset > count {
		offset = count
	}
	app.database.Model(&m).Offset(offset).Limit(limit).Select()

	for _, movie := range m {
		torrents := app.getMovieTorrents(movie.ID)
		if len(torrents) > 0 {
			var i float64
			for _, t := range torrents {
				i = i + t.NNM
			}
			movie.Torrent = torrents
			movie.NNM = round(i/float64(len(torrents)), 1)
			movies = append(movies, movie)
		}
	}
	data.Movies = movies
	data.Count = count
	data.Limit = len(movies)
	data.Offset = offset + data.Limit
	data.ImgDir = app.config.Web.ImgDir
	return data
}

func (app *application) getMovieTorrents(id int) []Torrent {
	var torrents []Torrent
	app.database.Model(&torrents).Where("movie_id = ?", id).Select()
	return torrents
}
