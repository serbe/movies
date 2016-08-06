package main

import "gopkg.in/pg.v4"

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
	ID          int       `sql:"id,pk"        json:"id"`
	Section     string    `sql:"section"      json:"section"`
	Name        string    `sql:"name"         json:"name"`
	EngName     string    `sql:"eng_name"     json:"eng_name"`
	Year        int       `sql:"year"         json:"year"`
	Genre       []string  `sql:"genre"        json:"genre"        pg:",array" `
	Country     []string  `sql:"country"      json:"country"      pg:",array"`
	RawCountry  string    `sql:"raw_country"  json:"raw_country"`
	Director    []string  `sql:"director"     json:"director"     pg:",array"`
	Producer    []string  `sql:"producer"     json:"producer"     pg:",array"`
	Actor       []string  `sql:"actor"        json:"actor"        pg:",array"`
	Description string    `sql:"description"  json:"description"`
	Age         string    `sql:"age"          json:"age"`
	ReleaseDate string    `sql:"release_date" json:"release_date"`
	RussianDate string    `sql:"russian_date" json:"russian_date"`
	Duration    string    `sql:"duration"     json:"duration"`
	Kinopoisk   float64   `sql:"kinopoisk"    json:"kinopoisk"`
	IMDb        float64   `sql:"imdb"         json:"imdb"`
	Poster      string    `sql:"poster"       json:"poster"`
	PosterURL   string    `sql:"poster_url"   json:"poster_url"`
	Torrent     []Torrent `sql:"-"            json:"torrent"`
	NNM         float64   `sql:"-"            json:"nnm"`
}

// Torrent all values
type Torrent struct {
	ID          int     `sql:"id,pk"             json:"id"`
	MovieID     int     `sql:"movie_id"          json:"movie_id"`
	DateCreate  string  `sql:"date_create"       json:"date_create"`
	Href        string  `sql:"href"              json:"href"`
	Torrent     string  `sql:"torrent"           json:"torrent"`
	Magnet      string  `sql:"magnet"            json:"magnet"`
	NNM         float64 `sql:"nnm"               json:"nnm"`
	Video       string  `sql:"video"             json:"video"`
	Quality     string  `sql:"quality"           json:"quality"`
	Resolution  string  `sql:"resolution"        json:"resolution"`
	Translation string  `sql:"translation"       json:"translation"`
	Size        int     `sql:"size"              json:"size"`
	Seeders     int     `sql:"seeders"           json:"seeders"`
	Leechers    int     `sql:"leechers"          json:"leechers"`
	// SubtitlesType string  `sql:"subtitles_type"`
	// Subtitles     string  `sql:"subtitles"`
	// Audio1        string  `sql:"audio1"`
	// Audio2        string  `sql:"audio2"`
	// Audio3        string  `sql:"audio3"`
}

type search struct {
	ID      int `sql:"id,pk"             json:"id"`
	MovieID int `sql:"movie_id"          json:"movie_id"`
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
		movies   []Movie
		count    int
		data     Data
		searches []search
	)

	count, _ = app.database.Model(&movies).Column("id").Count()
	if limit == 0 {
		limit = count
	}
	if offset > count {
		offset = count
	}
	//app.database.Model(&m).Order("id DESC").Offset(offset).Limit(limit).Select()
	// fast
	// EXPLAIN ANALYZE SELECT * FROM movies t1 JOIN (SELECT id FROM movies ORDER BY id LIMIT 10 OFFSET 150) as t2 ON t2.id = t1.id;
	app.database.Query(&searches, `SELECT max(id), movie_id from torrents group by movie_id order by max(id) desc LIMIT ? OFFSET ?;`, limit, offset)
	for _, s := range searches {
		movie := app.getMovieByID(s.MovieID)
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

func (app *application) getMovieByID(id int) Movie {
	var movie Movie
	app.database.Model(&movie).Where("id = ?", id).Select()
	return movie
}

func (app *application) getMovieTorrents(id int) []Torrent {
	var torrents []Torrent
	app.database.Model(&torrents).Where("movie_id = ?", id).Order("id DESC").Select()
	return torrents
}
