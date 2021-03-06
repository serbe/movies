package main

import (
	"log"
	"net/http"
	"sort"
	"strconv"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	if page == 0 {
		page = 1
	}
	movies, count, err := app.getMovies(page)
	if err != nil {
		log.Println("getAllMovies ", err)
		return
	}
	genres := []string{}
	years := []string{}
	for _, movie := range movies {
		year := strconv.Itoa(movie.Year)
		if !stringInSlice(years, year) {
			years = append(years, year)
		}
		filmGenres := movie.Genre
		for _, genre := range filmGenres {
			if !stringInSlice(genres, genre) {
				genres = append(genres, genre)
			}
		}
	}
	sort.Strings(genres)
	sort.Strings(years)
	cont := context{Title: "NNM-club RSS", Movies: movies, Genres: genres, Years: years, Count: count}
	_ = render(w, cont, "index")
}

// func (app *application) bindex(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query()
// 	page, _ := strconv.Atoi(query.Get("page"))
// 	if page == 0 {
// 		page = 1
// 	}
// 	movies, count, err := app.getMovies(page)
// 	if err != nil {
// 		log.Println("getAllMovies ", err)
// 		return
// 	}
// 	genres := []string{}
// 	years := []string{}
// 	for _, movie := range movies {
// 		year := strconv.Itoa(movie.Year)
// 		if stringInSlice(years, year) == false {
// 			years = append(years, year)
// 		}
// 		filmGenres := movie.Genre
// 		for _, genre := range filmGenres {
// 			if stringInSlice(genres, genre) == false {
// 				genres = append(genres, genre)
// 			}
// 		}
// 	}
// 	sort.Strings(genres)
// 	sort.Strings(years)
// 	cont := context{Title: "NNM-club RSS", Movies: movies, Genres: genres, Years: years, Count: count}
// 	renderBulma(w, cont, "bindex")
// }

func (app *application) search(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	year, _ := strconv.Atoi(query.Get("year"))
	actor := query.Get("actor")
	genre := query.Get("genre")
	page, _ := strconv.Atoi(query.Get("page"))
	if page == 0 {
		page = 1
	}
	director := query.Get("director")
	movies, count, err := app.searchMovies(page, year, actor, genre, director)
	if err != nil {
		log.Println("searchMovies ", 1, err)
		return
	}
	genres := []string{}
	years := []string{}
	for _, movie := range movies {
		year := strconv.Itoa(movie.Year)
		if !stringInSlice(years, year) {
			years = append(years, year)
		}
		filmGenres := movie.Genre
		for _, genre := range filmGenres {
			if !stringInSlice(genres, genre) {
				genres = append(genres, genre)
			}
		}
	}
	sort.Strings(genres)
	sort.Strings(years)
	cont := context{Title: "NNM-club RSS", Movies: movies, Genres: genres, Years: years, Count: count}
	_ = render(w, cont, "index")
}

// func (app *application) getOneMovieJSON(w http.ResponseWriter, r *http.Request) {
// 	t := time.Now()
// 	movies := app.getMovies(1, 0)
// 	data, err := json.Marshal(movies)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(data)
// 	app.printLog(t, r)
// }

// func (app *application) getMoviesJSON(w http.ResponseWriter, r *http.Request) {
// 	t := time.Now()
// 	query := r.URL.Query()
// 	limitStr := query.Get("limit")
// 	offsetStr := query.Get("offset")
// 	limit, _ := strconv.Atoi(limitStr)
// 	offset, _ := strconv.Atoi(offsetStr)
// 	movies := app.getMovies(limit, offset)
// 	data, err := json.Marshal(movies)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(data)
// 	app.printLog(t, r)
// }
