package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type config struct {
	Base struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
		Sslmode  string `json:"sslmode"`
	} `json:"postgresql"`
	Web struct {
		ImgDir string `json:"img_dir"`
		Port   string `json:"port"`
		Log    bool   `json:"log"`
	} `json:"web"`
}

func (app *application) getConfig() {
	c := config{}
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}
	app.cfg = c
}

func round(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}

func (app *application) printLog(startTime time.Time, r *http.Request) {
	if app.cfg.Web.Log {
		currentTime := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), currentTime.Sub(startTime))
	}
}

func stringInSlice(list []string, s string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}

func qs(s int) string {
	if s > 0 {
		return " AND"
	}
	return " WHERE"
}
