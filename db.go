package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Data - data to sent JSON
type Data struct {
	Offset int     `json:"Offset"`
	Count  int     `json:"Count"`
	Limit  int     `json:"Limit"`
	ImgDir string  `json:"ImgDir"`
	Movies []Movie `json:"Movies"`
}

// Movie all values
type Movie struct {
	ID          int64     `sql:"id,pk"        json:"id"`
	Section     string    `sql:"section"      json:"section"`
	Name        string    `sql:"name"         json:"name"`
	EngName     string    `sql:"eng_name"     json:"eng_name"`
	Year        int       `sql:"year"         json:"year"`
	Genre       []string  `sql:"genre"        json:"genre"`
	Country     []string  `sql:"country"      json:"country"`
	RawCountry  string    `sql:"raw_country"  json:"raw_country"`
	Director    []string  `sql:"director"     json:"director"`
	Producer    []string  `sql:"producer"     json:"producer"`
	Actor       []string  `sql:"actor"        json:"actor"`
	Description string    `sql:"description"  json:"description"`
	Age         string    `sql:"age"          json:"age"`
	ReleaseDate string    `sql:"release_date" json:"release_date"`
	RussianDate string    `sql:"russian_date" json:"russian_date"`
	Duration    string    `sql:"duration"     json:"duration"`
	Kinopoisk   float64   `sql:"kinopoisk"    json:"kinopoisk"`
	IMDb        float64   `sql:"imdb"         json:"imdb"`
	Poster      string    `sql:"poster"       json:"poster"`
	PosterURL   string    `sql:"poster_url"   json:"poster_url"`
	CreatedAt   time.Time `sql:"created_at"`
	UpdatedAt   time.Time `sql:"updated_at"`
	Torrent     []Torrent `sql:"-"            json:"torrent"`
	NNM         float64   `sql:"-"            json:"nnm"`
}

// Torrent all values
type Torrent struct {
	ID            int64     `sql:"id,pk"             json:"id"`
	MovieID       int64     `sql:"movie_id"          json:"movie_id"`
	DateCreate    string    `sql:"date_create"       json:"date_create"`
	Href          string    `sql:"href"              json:"href"`
	Torrent       string    `sql:"torrent"           json:"torrent"`
	Magnet        string    `sql:"magnet"            json:"magnet"`
	NNM           float64   `sql:"nnm"               json:"nnm"`
	SubtitlesType string    `sql:"subtitles_type"    json:"subtitles_type"`
	Subtitles     string    `sql:"subtitles"         json:"subtitles"`
	Video         string    `sql:"video"             json:"video"`
	Quality       string    `sql:"quality"           json:"quality"`
	Resolution    string    `sql:"resolution"        json:"resolution"`
	Audio1        string    `sql:"audio1"            json:"audio1"`
	Audio2        string    `sql:"audio2"            json:"audio2"`
	Audio3        string    `sql:"audio3"            json:"audio3"`
	Translation   string    `sql:"translation"       json:"translation"`
	Size          int       `sql:"size"              json:"size"`
	Seeders       int       `sql:"seeders"           json:"seeders"`
	Leechers      int       `sql:"leechers"          json:"leechers"`
	CreatedAt     time.Time `sql:"created_at"`
	UpdatedAt     time.Time `sql:"updated_at"`
}

type search struct {
	ID      int64 `sql:"id,pk"             json:"id"`
	MovieID int64 `sql:"movie_id"          json:"movie_id"`
}

func (app *application) initDB() {
	options := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		app.cfg.Base.Dbname,
		app.cfg.Base.User,
		app.cfg.Base.Password,
		app.cfg.Base.Sslmode,
	)
	db, err := sql.Open("postgres", options)
	if err != nil {
		log.Fatal(err)
	}
	app.db = db
}

func (app *application) getMovies(limit int, offset int) Data {
	var (
		movies []Movie
		count  int
		data   Data
	)

	_ = app.db.QueryRow("SELECT count(*) FROM movies;").Scan(&count)
	if limit == 0 {
		limit = count
	}
	if offset > count {
		offset = count
	}
	//app.database.Model(&m).Order("id DESC").Offset(offset).Limit(limit).Select()
	// fast
	// EXPLAIN ANALYZE SELECT * FROM movies t1 JOIN (SELECT id FROM movies ORDER BY id LIMIT 10 OFFSET 150) as t2 ON t2.id = t1.id;
	rows, err := app.db.Query(`SELECT max(id), movie_id FROM torrents GROUP BY movie_id ORDER BY max(id) desc LIMIT $1 OFFSET $2;`, limit, offset)
	if err != nil {
		return data
	}
	resultMovies, err := scanSearchs(rows)
	if err != nil {
		return data
	}
	for _, s := range resultMovies {
		movie, _ := app.getMovieByID(s.MovieID)
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
	data.Movies = movies
	data.Count = count
	data.Limit = len(movies)
	data.Offset = offset + data.Limit
	data.ImgDir = app.cfg.Web.ImgDir
	return data
}

func (app *application) getMovieByID(id int64) (Movie, error) {
	row := app.db.QueryRow("Select * from FROM movies WHERE id = $1", id)
	return scanMovie(row)
}

func (app *application) getMovieTorrents(id int64) ([]Torrent, error) {
	rows, err := app.db.Query("SELECT * FROM torrents WHERE movie_id = $1 ORDER BY seeders DESC", id)
	if err != nil {
		return nil, err
	}
	return scanTorrents(rows)
}
