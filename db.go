package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
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
	ID          int64          `sql:"id,pk"        json:"id"`
	Section     string         `sql:"section"      json:"section"`
	Name        string         `sql:"name"         json:"name"`
	EngName     string         `sql:"eng_name"     json:"eng_name"`
	Year        int            `sql:"year"         json:"year"`
	Genre       pq.StringArray `sql:"genre"        json:"genre"`
	Country     pq.StringArray `sql:"country"      json:"country"`
	RawCountry  string         `sql:"raw_country"  json:"raw_country"`
	Director    pq.StringArray `sql:"director"     json:"director"`
	Producer    pq.StringArray `sql:"producer"     json:"producer"`
	Actor       pq.StringArray `sql:"actor"        json:"actor"`
	Description string         `sql:"description"  json:"description"`
	Age         string         `sql:"age"          json:"age"`
	ReleaseDate string         `sql:"release_date" json:"release_date"`
	RussianDate string         `sql:"russian_date" json:"russian_date"`
	Duration    string         `sql:"duration"     json:"duration"`
	Kinopoisk   float64        `sql:"kinopoisk"    json:"kinopoisk"`
	IMDb        float64        `sql:"imdb"         json:"imdb"`
	Poster      string         `sql:"poster"       json:"poster"`
	PosterURL   string         `sql:"poster_url"   json:"poster_url"`
	CreatedAt   pq.NullTime    `sql:"created_at"`
	UpdatedAt   pq.NullTime    `sql:"updated_at"`
	Torrent     []Torrent      `sql:"-"            json:"torrent"`
	NNM         float64        `sql:"-"            json:"nnm"`
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
		app.cfg.Base.User,
		app.cfg.Base.Password,
		app.cfg.Base.Dbname,
		app.cfg.Base.Sslmode,
	)
	db, err := sql.Open("postgres", options)
	if err != nil {
		log.Fatal(err)
	}
	app.db = db
}

func (app *application) getMovies(page int) ([]Movie, int64, error) {
	var (
		movies []Movie
		count  int64
	)

	_ = app.db.QueryRow("SELECT count(*) FROM movies;").Scan(&count)
	// EXPLAIN ANALYZE SELECT * FROM movies t1 JOIN (SELECT id FROM movies ORDER BY id LIMIT 10 OFFSET 150) as t2 ON t2.id = t1.id;
	rows, err := app.db.Query(`SELECT max(id), movie_id FROM torrents GROUP BY movie_id ORDER BY max(id) desc LIMIT $1 OFFSET $2;`, 50, (page-1)*50)
	if err != nil {
		log.Println("Query search ", err)
		return nil, 0, err
	}
	searches, err := scanSearchs(rows)
	if err != nil {
		log.Println("scanSearchs ", err)
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

func (app *application) getAllMovies() ([]Movie, error) {
	var movies []Movie
	rows, err := app.db.Query(`SELECT max(id), movie_id FROM torrents GROUP BY movie_id ORDER BY max(id) desc`)
	if err != nil {
		DBError := err.(*pq.Error) // for Postgres DB driver
		fmt.Println("SQL ERROR!")
		fmt.Printf("%#v\n", DBError)
		log.Println("Query search ", err)
		return nil, err
	}
	searches, err := scanSearchs(rows)
	if err != nil {
		log.Println("scanSearchs ", err)
		return nil, err
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
		} else {
			log.Println("getMovieTorrents ", movie.ID, err)
		}
	}
	return movies, nil
}

func (app *application) getMovieByID(id int64) (Movie, error) {
	row := app.db.QueryRow("Select * FROM movies WHERE id = $1", id)
	return scanMovie(row)
}

func (app *application) getMovieTorrents(id int64) ([]Torrent, error) {
	rows, err := app.db.Query("SELECT * FROM torrents WHERE movie_id = $1 ORDER BY seeders DESC", id)
	if err != nil {
		return nil, err
	}
	return scanTorrents(rows)
}
