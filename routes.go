package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"io/ioutil"
)

func root(w http.ResponseWriter, r *http.Request)  {
	data, err := ioutil.ReadFile("./templates/root.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (app *application) getOneMovieJSON(w http.ResponseWriter, r *http.Request) {
	movies := app.getMovies(1, 0)
	data, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (app *application) getMoviesJSON(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	movies := app.getMovies(limit, offset)
	data, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
