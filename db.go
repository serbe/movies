package main

import (
	"log"

	pg "gopkg.in/pg.v5"
)

// Data - data to sent JSON
type Data struct {
	Offset int      `json:"Offset"`
	Count  int      `json:"Count"`
	Limit  int      `json:"Limit"`
	ImgDir string   `json:"ImgDir"`
	Genges []string `json:"Genres"`
	Years  []int64  `json:"Years"`
	Movies []Movie  `json:"Movies"`
}

// Movie all values
type Movie struct {
	ID          int64       `sql:"id,pk"        json:"id"`
	Section     string      `sql:"section"      json:"section"`
	Name        string      `sql:"name"         json:"name"`
	EngName     string      `sql:"eng_name"     json:"eng_name"`
	Year        int         `sql:"year"         json:"year"`
	Genre       []string    `sql:"genre"        json:"genre"         pg:",array"`
	Country     []string    `sql:"country"      json:"country"       pg:",array"`
	RawCountry  string      `sql:"raw_country"  json:"raw_country"`
	Director    []string    `sql:"director"     json:"director"      pg:",array"`
	Producer    []string    `sql:"producer"     json:"producer"      pg:",array"`
	Actor       []string    `sql:"actor"        json:"actor"         pg:",array"`
	Description string      `sql:"description"  json:"description"`
	Age         string      `sql:"age"          json:"age"`
	ReleaseDate string      `sql:"release_date" json:"release_date"`
	RussianDate string      `sql:"russian_date" json:"russian_date"`
	Duration    string      `sql:"duration"     json:"duration"`
	Kinopoisk   float64     `sql:"kinopoisk"    json:"kinopoisk"`
	IMDb        float64     `sql:"imdb"         json:"imdb"`
	Poster      string      `sql:"poster"       json:"poster"`
	PosterURL   string      `sql:"poster_url"   json:"poster_url"`
	CreatedAt   pg.NullTime `sql:"created_at"`
	UpdatedAt   pg.NullTime `sql:"updated_at"`
	Torrent     []Torrent   `sql:"-"            json:"torrent"`
	NNM         float64     `sql:"-"            json:"nnm"`
}

// Torrent all values
type Torrent struct {
	ID            int64       `sql:"id,pk"             json:"id"`
	MovieID       int64       `sql:"movie_id"          json:"movie_id"`
	DateCreate    string      `sql:"date_create"       json:"date_create"`
	Href          string      `sql:"href"              json:"href"`
	Torrent       string      `sql:"torrent"           json:"torrent"`
	Magnet        string      `sql:"magnet"            json:"magnet"`
	NNM           float64     `sql:"nnm"               json:"nnm"`
	SubtitlesType string      `sql:"subtitles_type"    json:"subtitles_type"`
	Subtitles     string      `sql:"subtitles"         json:"subtitles"`
	Video         string      `sql:"video"             json:"video"`
	Quality       string      `sql:"quality"           json:"quality"`
	Resolution    string      `sql:"resolution"        json:"resolution"`
	Audio1        string      `sql:"audio1"            json:"audio1"`
	Audio2        string      `sql:"audio2"            json:"audio2"`
	Audio3        string      `sql:"audio3"            json:"audio3"`
	Translation   string      `sql:"translation"       json:"translation"`
	Size          int         `sql:"size"              json:"size"`
	Seeders       int         `sql:"seeders"           json:"seeders"`
	Leechers      int         `sql:"leechers"          json:"leechers"`
	CreatedAt     pg.NullTime `sql:"created_at"`
	UpdatedAt     pg.NullTime `sql:"updated_at"`
}

type search struct {
	ID      int64 `sql:"id,pk"             json:"id"`
	MovieID int64 `sql:"movie_id"          json:"movie_id"`
}

func (app *application) initDB() {
	options := &pg.Options{
		User:     app.cfg.Base.User,
		Password: app.cfg.Base.Password,
		Database: app.cfg.Base.Dbname,
		// app.cfg.Base.Sslmode,
	}
	db := pg.Connect(options)
	app.db = db
}

func (app *application) getMovies(page int) ([]Movie, int, error) {
	var (
		movies   []Movie
		searches []search
	)

	count, err := app.db.Model(&Movie{}).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = app.db.Query(&searches, `SELECT max(t.id), t.movie_id FROM torrents AS t GROUP BY movie_id ORDER BY max(id) desc LIMIT ? OFFSET ?;`, 50, (page-1)*50)
	if err != nil {
		log.Println("Query search ", err)
		return nil, 0, err
	}
	for _, s := range searches {
		movie, err := app.getMovieByID(s.MovieID)
		if err != nil {
			log.Println("getMovieByID ", s.MovieID, err)
		}
		torrents, err := app.getMovieTorrents(movie.ID)
		if err == nil && len(torrents) > 0 {
			var i float64
			for _, t := range torrents {
				i = i + t.NNM
			}
			movie.Torrent = torrents
			movie.NNM = round(i/float64(len(torrents)), 1)
			movies = append(movies, movie)
		}
	}
	return movies, count, nil
}

func (app *application) getMovieByID(id int64) (Movie, error) {
	var movie Movie
	_, err := app.db.QueryOne(&movie, "Select * FROM movies WHERE id = ?", id)
	return movie, err
}

func (app *application) getMovieTorrents(id int64) ([]Torrent, error) {
	var torrents []Torrent
	_, err := app.db.Query(&torrents, "SELECT * FROM torrents WHERE movie_id = ? ORDER BY seeders DESC", id)
	return torrents, err
}
