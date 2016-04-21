package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
	Base struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
		Sslmode  bool   `json:"sslmode"`
	} `json:"postgresql"`
	Web struct {
		Port string `json:"port"`
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
	app.config = c
	log.Println(app.config)
}

func round(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}